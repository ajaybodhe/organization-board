package models

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (login *Login) Valid() bool {
	if "" == login.Email || "" == login.Password {
		return false
	}
	return true
}
