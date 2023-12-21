package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	client := shell.NewShell("/ip4/127.0.0.1/tcp/5001")
	file, _ := os.Open("/tmp/aaa.txt")
	cid, err := client.Add(file, shell.Pin(true))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cid)
}

func TestGenerateRandomString(t *testing.T) {
	length := 20
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(bytes)[:length])
}
