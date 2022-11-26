package main

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

const password = "wnKisl81Y"
const salt = "abcdef"

func main() {
	salt := []byte(salt)
	// func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error)
	// AES-256（32バイトの鍵を必要とする）の場合
	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(dk))
}
