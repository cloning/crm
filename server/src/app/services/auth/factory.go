package auth

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
