package auth

import (
	"crypto/sha512"
	"encoding/base64"
)

func hashPassword(password string) string {
	d := sha512.New()
	d.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(d.Sum(nil))
}
