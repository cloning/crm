package service

import (
	"../../../../src/app/services/auth"
	"os"
	"testing"
	"time"
)

const (
	AUTH_FILE_PATH = "./mock/.admin_accounts"
	DATA_DIR       = "./data"
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

func createAuthService(t *testing.T, callback AuthServiceCallback) {
	if err := os.Mkdir(DATA_DIR, 0777); os.IsNotExist(err) {
		t.Error(err)
	}
	defer func() {
		err := os.RemoveAll(DATA_DIR)
		if err != nil {
			t.Error(err)
		}
	}()

	tokenRepository, err := auth.NewFileTokenRepository(DATA_DIR)
	if err != nil {
		t.Errorf("Coudn't create token repository %v", err)
	}
	authService, err := auth.NewAuthService(AUTH_FILE_PATH, tokenRepository)

	if err != nil {
		t.Errorf("Coudn't create auth service %v", err)
	}
	callback(authService)
}
