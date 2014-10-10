package auth

import (
	"time"
)

type TokenRepository interface {
	Get(string) (*Token, error)
	Save(*Token) error
}

type AuthService struct {
	adminUsers      map[string]string
	tokenRepository TokenRepository
}

func (this *AuthService) Validate(userName, password string) *Token {

	hashed := hashPassword(password)

	if this.adminUsers[userName] == hashed {
		token := &Token{userName, generateTokenKey(), time.Now().Add(10 * time.Hour)}
		// TODO: Handle error
		this.tokenRepository.Save(token)
		return token
	}

	return nil
}

func (this *AuthService) ValidateToken(key string) bool {
	//TODO: Handle error
	token, _ := this.tokenRepository.Get(key)
	return token != nil
}
