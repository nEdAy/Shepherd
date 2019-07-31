package scrypt

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
)

func GetScryptPasswordBase64(password string, salt string) string {
	dkPassword, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	passwordBase64 := base64.StdEncoding.EncodeToString(dkPassword)
	return passwordBase64
}
