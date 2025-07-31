package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

// command to generate a 32-byte key and pepper:
// openssl rand -hex 32

var (
	AES_SECRET string = "AES_SECRET"
	AES_PEPPER string = "AES_PEPPER"
)

type Crypto interface {
	Encrypt(plainText string) ([]byte, error)
	Decrypt(encryptedTextBytes []byte) (string, error)
	Hash(plainText string) []byte
}

type AES256Crypto struct {
	key       []byte
	pepper    []byte
	nonceSize int
}

func NewAES256Crypto() Crypto {
	key := os.Getenv(AES_SECRET)
	pepper := os.Getenv(AES_PEPPER)

	bKey, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}
	bPepper, err := hex.DecodeString(pepper)
	if err != nil {
		panic(err)
	}

	if len(bKey) != 32 || len(bPepper) != 32 {
		panic("Key and Pepper must be 32 bytes for AES-256")
	}

	return &AES256Crypto{key: bKey, pepper: bPepper, nonceSize: 12}
}

func (a *AES256Crypto) Encrypt(plainText string) ([]byte, error) {
	// Create a new AES cipher block from the key.
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, a.nonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// GCM is an authenticated encryption mode that provides both confidentiality and authenticity.
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(plainText), nil)
	fullCiphertext := append(nonce, ciphertext...)

	return fullCiphertext, nil
}

func (a *AES256Crypto) Decrypt(encryptedTextBytes []byte) (string, error) {
	if len(encryptedTextBytes) < a.nonceSize {
		return "", errors.New("ciphertext too short")
	}

	// Create a new AES cipher block from the key.
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	// GCM is an authenticated encryption mode that provides both confidentiality and authenticity.
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce, ciphertext := encryptedTextBytes[:a.nonceSize], encryptedTextBytes[a.nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (a *AES256Crypto) Hash(plainText string) []byte {
	// Create a new HMAC object with SHA256
	h := hmac.New(sha256.New, a.pepper)
	h.Write([]byte(plainText))

	return h.Sum(nil)
}
