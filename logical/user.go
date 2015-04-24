package logical

import (
	"github.com/perfectworks/goboard/storage"
	"github.com/perfectworks/goboard/web/vo"
)

// SaveUser will create or update a user
func SaveUser(user *vo.User, ctx *Context) (err error) {
	mUser := user.Model()

	err = mUser.Save(ctx.DbMap)
	if err != nil {
		return err
	}

	newUser := vo.NewUser(mUser)

	*user = *newUser

	return
}

// GetUserByEmail retrieve user by email
func GetUserByEmail(email string, ctx *Context) (user *vo.User, err error) {
	mUser, err := storage.GetUserByEmail(email, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	return vo.NewUser(mUser), nil
}

// GetUserBySalt retrieve user by salt
func GetUserBySalt(salt string, ctx *Context) (user *vo.User, err error) {
	mUser, err := storage.GetUserBySalt(salt, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	return vo.NewUser(mUser), nil
}

// InitUser will retrieve user by email, or create user if not existed
func InitUser(email string, ctx *Context) (user *vo.User, err error) {
	user, err = GetUserByEmail(email, ctx)
	if err != nil {
		mUser := &storage.User{Email: email}
		err = mUser.Save(ctx.DbMap)
		if err != nil {
			return nil, err
		}

		user = vo.NewUser(mUser)
	}

	return user, nil
}

// QueryUser will retrieve users by email
func QueryUser(email string, ctx *Context) (users []vo.User, err error) {
	mUsers, err := storage.QueryUser(email, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	users = make([]vo.User, len(mUsers))
	for i, v := range mUsers {
		users[i] = *vo.NewUser(&v)
	}

	return
}
