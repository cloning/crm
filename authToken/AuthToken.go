package authToken

import (
	"time"
)

type AuthToken struct {
	Code    string    `json:"code"`
	UserId  string    `json:"user"`
	Created time.Time `json:"created"`
	Expires time.Time `json:"expires"`
}

func Create(userId string) *AuthToken {
	return &AuthToken{
		Code:    "asd123",
		Created: time.Now(),
		Expires: time.Now().Add(24 * time.Hour),
		UserId:  userId,
	}
}

func Invalidate(code, userId string) {

}
