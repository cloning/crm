package auth

import (
	"time"
)

type Token struct {
	UserName string
	Key      string
	Expires  time.Time
}
