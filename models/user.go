package models

import (
	"github.com/yuantiku/rndstr"
	"gopkg.in/gorp.v1"
)

type User struct {
	Id    int    `db:"id"`
	Email string `db:"email"`
	Salt  string `db:"salt"`
}

func initUserTable(dbmap *gorp.DbMap) {
	userTable := dbmap.AddTableWithName(User{}, "users")
	userTable.SetKeys(true, "id")
	userTable.ColMap("email").SetUnique(true)
}

func GetUserById(id int, dbmap *gorp.DbMap) *User {
	user := &User{}
	err := dbmap.SelectOne(&user, "select * from users where id = ?", id)
	if err != nil {
		return nil
	}

	return user
}

func GetAllUser(dbmap *gorp.DbMap) []User {
	var users []User
	_, err := dbmap.Select(&users, "select * from users")
	panicIf(err)

	return users
}

func QueryUserByEmail(email string, dbmap *gorp.DbMap) []User {
	var users []User
	_, err := dbmap.Select(&users, "select * from users where email like ?", "%"+email+"%")
	panicIf(err)

	return users
}

func InitUser(email string, dbmap *gorp.DbMap) *User {
	user := &User{}
	err := dbmap.SelectOne(&user, "select * from users where email = ?", email)
	if err != nil {
		user.Email = email
		user.Save(dbmap)
	}

	return user
}

func (u *User) Save(dbmap *gorp.DbMap) {
	if len(u.Salt) == 0 {
		u.Salt = rndstr.Gen(32)
	}

	err := dbmap.Insert(u)
	panicIf(err)
}

func (u *User) Update(dbmap *gorp.DbMap) {
	count, err := dbmap.Update(u)
	panicIf(err)

	if count == 0 {
		panic("save user failed")
	}
}
