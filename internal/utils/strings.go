package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	mathrand "math/rand"
	"time"
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

func GetOutTradeNo() string {
	currentTime := time.Now()
	datePart := currentTime.Format("20060102150405")
	serialNumber := "0001"
	randomNumber := mathrand.New(mathrand.NewSource(time.Now().UnixNano())).Int31n(100000000)
	return fmt.Sprintf("cycle%s%s%08d", datePart, serialNumber, randomNumber)
}
