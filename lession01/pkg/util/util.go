package util

import (
	"encoding/base64"
)

func Encrypt(plainText string) string {
	return base64.StdEncoding.EncodeToString([]byte(plainText))
}

func Decrypt(encryptText string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(encryptText)
	return string(bytes), err
}
