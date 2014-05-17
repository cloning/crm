package user

// Roles
type Role string

const (
	Admin    Role = "admin"
	Supplier Role = "supplier"
	Advisor  Role = "advisor"
	Customer Role = "customer"
)

type User struct {
	// Public
	Id              string `json:"id"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	TelephoneNumber string `json:"telephoneNumber"`
	Roles           []Role `json:"roles"`

	// Private
	password string
}

func FindFromCredentials(email, password string) *User {
	return &User{
		"123",
		"Test",
		"Test",
		email,
		"07037724444",
		make([]Role, 0),
		"nil",
	}
}
