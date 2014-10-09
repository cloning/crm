package auth

import (
	"time"
)

type AuthService struct {
	adminUsers map[string]string
}

type Token struct {
	Token   string
	Expires time.Time
}

func NewAuthService(accountsFile string) (*AuthService, error) {
	users, err := parseAuthFile(accountsFile)

	if err != nil {
		return nil, err
	}

	return &AuthService{users}, nil
}

func (this *AuthService) Validate(userName, password string) *Token {
	hashed := hashPassword(password)
	if this.adminUsers[userName] == hashed {
		return &Token{"Test", time.Now().Add(10 * time.Hour)}
	}
	return nil
}
