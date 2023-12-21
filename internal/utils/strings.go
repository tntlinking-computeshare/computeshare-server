package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func ToAddress(str string) *string {
	ptr := &str
	return ptr
}

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

func StringKeyDesensitization(key string) string {
	if len(key) > 4 {
		return key[:2] + "******" + key[len(key)-2:]
	}
	return key
}
