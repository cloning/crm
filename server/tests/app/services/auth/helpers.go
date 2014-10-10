package service

import (
	"../../../../src/app/services/auth"
	"os"
	"testing"
	"time"
)

type AuthServiceCallback func(*auth.AuthService)

func verifyToken(token *auth.Token, t *testing.T) {
	if token == nil {
		t.Log("Token was nil")
		t.Fail()
	}

	if token.Key == "" {
		t.Log("Token key was nil")
		t.Fail()
	}

	now := time.Now()

	if token.Expires.Before(now) || token.Expires.Equal(now) {
		t.Log("Token already expired")
		t.Fail()
	}
}

func createAuthService(t *testing.T, accountsFile string, callback AuthServiceCallback) {
	if err := os.Mkdir("./data", 0777); os.IsNotExist(err) {
		t.Error(err)
	}
	defer func() {
		err := os.RemoveAll("./data")
		if err != nil {
			t.Error(err)
		}
	}()

	tokenRepository := auth.NewFileTokenRepository("./data")
	authService, err := auth.NewAuthService(accountsFile, tokenRepository)

	if err != nil {
		t.Errorf("Coudn't create auth service %v", err)
	}
	callback(authService)
}
