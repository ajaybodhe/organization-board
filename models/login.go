package models

// Login : login credentials of user
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Valid : check if Login structure is valid
func (login *Login) Valid() bool {
	if "" == login.Email || "" == login.Password {
		return false
	}
	return true
}
