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
	users, err := parseAccountsFile(accountsFile)

	if err != nil {
		return nil, err
	}

	authService := &AuthService{
		adminUsers: users,
	}

	return authService, nil
}

func (this *AuthService) Validate(userName, password string) *Token {

	hashed := hashPassword(password)

	// First check if administrator
	if this.adminUsers[userName] == hashed {
		return &Token{"Test", time.Now().Add(10 * time.Hour)}
	}

	// TODO: Other user types here
	return nil
}
