package storage

import (
	"errors"

	"github.com/yuantiku/rndstr"
	"gopkg.in/gorp.v1"
)

// User is basic user object
type User struct {
	ID    int    `db:"id"`
	Email string `db:"email"`
	Salt  string `db:"salt"`
}

func initUserTable(dbmap *gorp.DbMap) {
	userTable := dbmap.AddTableWithName(User{}, "users")
	userTable.SetKeys(true, "id")
	userTable.ColMap("email").SetUnique(true)
}

// GetUserByID will retrieve a user by specified user id
func GetUserByID(id int, dbmap *gorp.DbMap) (user *User) {
	err := dbmap.SelectOne(&user, "select * from users where id = ?", id)
	if err != nil {
		return nil
	}

	return
}

// GetAllUser will retrieve all users
func GetAllUser(dbmap *gorp.DbMap) (users []User, err error) {
	_, err = dbmap.Select(&users, "select * from users")
	return
}

// QueryUserByEmail will retrieve users match specified email address snippet
func QueryUserByEmail(email string, dbmap *gorp.DbMap) (users []User, err error) {
	_, err = dbmap.Select(&users, "select * from users where email like ?", "%"+email+"%")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// InitUser try to retrieve user by specified email, and create a new one if no user found
func InitUser(email string, dbmap *gorp.DbMap) (user *User, err error) {
	user = &User{}
	err = dbmap.SelectOne(&user, "select * from users where email = ?", email)
	if err == nil {
		return
	}

	user.Email = email
	err = user.Save(dbmap)
	if err != nil {
		return nil, err
	}

	return
}

// Save will insert a new user record to database
func (u *User) Save(dbmap *gorp.DbMap) (err error) {
	if len(u.Salt) == 0 {
		u.ResetSalt()
	}

	var affectCount int64
	if u.ID > 0 {
		affectCount, err = dbmap.Update(u)
	} else {
		err = dbmap.Insert(u)
		affectCount = 1
	}

	if err != nil {
		return err
	}

	if affectCount == 0 {
		return errors.New("failed to save user, affectCount = 0")
	}

	id := u.ID
	rowCount, err := dbmap.SelectInt("select count(*) as count from users where id < ?", id)
	if err != nil {
		return err
	}

	if rowCount != 0 {
		return nil
	}

	roles, err := QueryRoleByScope(RoleGlobal, dbmap)
	if err != nil {
		return err
	}

	if len(roles) == 0 {
		return errors.New("Administrator role is not existed")
	}

	admin := roles[0]
	userRole := &UserRole{UserID: id, RoleID: admin.ID, ProjectID: ProjectNone}
	err = userRole.Save(dbmap)
	if err != nil {
		return err
	}

	return
}

// ResetSalt will regenerate a valid salt for user
func (u *User) ResetSalt() {
	u.Salt = rndstr.Gen(32)
}
