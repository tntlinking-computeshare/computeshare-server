package service

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	shell "github.com/ipfs/go-ipfs-api"
	"os"
	"testing"
	"time"
)

const hmacSampleSecret = "some-secret-key-for-forntend"

func TestAdd(t *testing.T) {
	client := shell.NewShell("/ip4/127.0.0.1/tcp/5001")
	file, _ := os.Open("/tmp/aaa.txt")
	cid, err := client.Add(file, shell.Pin(true))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cid)
}

func TestJWT(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIxNzU1OWRlZi1jNTE0LTQxZmQtYjA0NC1iNTliODFmMmZiNmIiLCJleHAiOjE3MDMwNDI4NzZ9.z99Trv7jJ0aRNIijyyaFRU0fygFKrPwEcNgYUii5N4s"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSampleSecret), nil
	})
	if err != nil {
		panic(err)
	}

	if claim, ok := token.Claims.(jwt.MapClaims); token.Valid && ok {
		fmt.Println(claim["UserID"])
	} else {
		fmt.Println(token)
	}
}

func TestJwt2(t *testing.T) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))

	fmt.Println(tokenString, err)

	token2, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSampleSecret), nil
	})

	if claims, ok := token2.Claims.(jwt.MapClaims); ok && token2.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}
