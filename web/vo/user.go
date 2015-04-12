package vo

import "github.com/yuantiku/goboard/storage"

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Salt  string `json:"salt"`
}

func (u *User) Model() (user *storage.User, err error) {
	user = &storage.User{
		ID:    u.ID,
		Email: u.Email,
		Salt:  u.Salt}

	return
}
