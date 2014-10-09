package service

import (
	"../../../../src/app/services/auth"
	"testing"
	"time"
)

func TestAdminAccounts(t *testing.T) {
	accountsFile := "./mock/.admin_accounts"
	authService, err := auth.NewAuthService(accountsFile)

	if err != nil {
		t.Errorf("Coudn't create auth service %v", err)
	}

	userName, password := "admin", "pwd"

	token := authService.Validate(userName, password)

	if token == nil {
		t.Log("Token was nil")
		t.Fail()
	}

	if token.Token == "" {
		t.Log("Token key was nil")
		t.Fail()
	}

	now := time.Now()

	if token.Expires.Before(now) || token.Expires.Equal(now) {
		t.Log("Token already expired")
		t.Fail()
	}
}
