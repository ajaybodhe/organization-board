package models

// User : user details in system
type User struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

// String : return string version of User
func (user *User) String() string {
	return Stringify(user)
}
