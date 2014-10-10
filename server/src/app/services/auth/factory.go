package auth

import (
	"os"
)

func NewAuthService(accountsFile string, tokenRepository TokenRepository) (*AuthService, error) {

	users, err := parseAccountsFile(accountsFile)

	if err != nil {
		return nil, err
	}

	authService := &AuthService{
		adminUsers:      users,
		tokenRepository: tokenRepository,
	}

	return authService, nil
}

func NewFileTokenRepository(dataDirectory string) (TokenRepository, error) {
	if _, err := os.Stat(dataDirectory); err != nil {
		return nil, err
	}
	return TokenRepository(&FileTokenRepository{dataDirectory}), nil
}
