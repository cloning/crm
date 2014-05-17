package core

import (
	"github.com/cloning/crm/authToken"
	"github.com/cloning/crm/user"
)

func Auth_Login(email, password string) *authToken.AuthToken {
	u := user.FindFromCredentials(email, password)

	if u == nil {
		return nil
	}

	return authToken.Create(u.Id)
}

func Auth_Logout(code, userId string) {
	authToken.Invalidate(code, userId)
}
