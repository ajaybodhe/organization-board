package models

// User : user details in system
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// String : return string version of User
func (user *User) String() string {
	return Stringify(user)
}
