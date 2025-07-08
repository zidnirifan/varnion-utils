package action

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func PATCH(url string, payload []byte) ([]byte, int, error) {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, http.StatusConflict, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return body, resp.StatusCode, nil
}
