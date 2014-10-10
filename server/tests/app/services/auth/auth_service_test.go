package service

import (
	"../../../../src/app/services/auth"
	"testing"
)

const (
	AUTH_FILE_PATH = "./mock/.admin_accounts"
)

/*
	Asserts that valid accounts give valid tokens
*/
func TestAdminAccounts(t *testing.T) {
	createAuthService(t, AUTH_FILE_PATH, func(authService *auth.AuthService) {

		t1 := authService.Validate("admin", "pwd")
		t2 := authService.Validate("otheradmin", "pwd")
		t3 := authService.Validate("non-existing", "non-existing")

		if t1.Key == t2.Key {
			t.Log("Token keys are not unique")
			t.Fail()
		}

		verifyToken(t1, t)
		verifyToken(t2, t)

		if t3 != nil {
			t.Log("User should not be valid")
			t.Fail()
		}
	})
}

func TestTokenValidate(t *testing.T) {
	createAuthService(t, AUTH_FILE_PATH, func(authService *auth.AuthService) {

		token := authService.Validate("admin", "pwd")

		if authService.ValidateToken(token.Key) == false {
			t.Log("Token should be valid")
			t.Fail()
		}

		token.Key = "invalid"

		if authService.ValidateToken(token.Key) == true {
			t.Log("Token shouldn't be valid")
			t.Fail()
		}
	})
}

/*
	Asserts that a token is persisted across instances of services
*/
func TestTokenPersistance(t *testing.T) {
	createAuthService(t, AUTH_FILE_PATH, func(authService1 *auth.AuthService) {
		createAuthService(t, AUTH_FILE_PATH, func(authService2 *auth.AuthService) {
			token := authService1.Validate("admin", "pwd")

			verifyToken(token, t)

			validInFirst := authService1.ValidateToken(token.Key)
			validInSecond := authService2.ValidateToken(token.Key)

			if validInFirst == false {
				t.Log("Token was not valid in first service instance")
				t.Fail()
			}

			if validInSecond == false {
				t.Log("Token was not valid in first service instance")
				t.Fail()
			}
		})
	})
}
