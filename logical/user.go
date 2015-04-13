package logical

import "github.com/yuantiku/goboard/web/vo"

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
