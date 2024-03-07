package biz

import (
	"fmt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"testing"
)

func TestJWT(t *testing.T) {
	ApiKey := "some-secret-key-for-forntend"
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIxNzU1OWRlZi1jNTE0LTQxZmQtYjA0NC1iNTliODFmMmZiNmIiLCJleHAiOjE3MDgzOTU2NDh9.GgWrHQ6dnBMAw7lf8WS_WUH8scjMWfORaQrHa_6rhfI"
	token, err := jwt2.Parse(tokenString, func(token *jwt2.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt2.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(ApiKey), nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("token:", token)
}
