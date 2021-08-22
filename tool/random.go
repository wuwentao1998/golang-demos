package tool

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomToken(size uint) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
