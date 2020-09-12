package models

type User struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

func (user *User) String() string {
	return Stringify(user)
}
