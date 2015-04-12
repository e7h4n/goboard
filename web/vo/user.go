package vo

import "github.com/yuantiku/goboard/storage"

// User is view object for storage.User
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Salt  string `json:"salt"`
}

// Model convert vo to storage model
func (u *User) Model() (user *storage.User) {
	return &storage.User{
		ID:    u.ID,
		Email: u.Email,
		Salt:  u.Salt}
}

// NewUser convert storage model to vo
func NewUser(u *storage.User) (user *User) {
	return &User{
		ID:    u.ID,
		Email: u.Email,
		Salt:  u.Salt}
}
