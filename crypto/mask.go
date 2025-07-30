package crypto

import "strings"

// Mask the string by replacing all characters with '*'
// except for the first and last characters.
func Mask(s string) string {
	if len(s) <= 2 {
		return "**"
	}
	runes := []rune(s)
	for i := 1; i < len(runes)-1; i++ {
		runes[i] = '*'
	}

	return string(runes)
}

func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return Mask(email)
	}
	localPart := parts[0]
	domainPart := parts[1]

	maskedLocal := localPart[:1] + strings.Repeat("*", len(localPart)-1)
	maskedDomain := strings.Repeat("*", len(domainPart)-1) + domainPart[len(domainPart)-1:]

	return maskedLocal + "@" + maskedDomain
}
