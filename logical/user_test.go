package logical

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuantiku/goboard/storage"
	"github.com/yuantiku/goboard/web/vo"
)

var salt string

func TestSaveUser(t *testing.T) {
	ctx := initTest(true)

	user := &vo.User{Email: "admin@fenbi.com"}
	err := SaveUser(user, ctx)
	checkTestErr(err)

	assert.Equal(t, 1, user.ID)
	mUser, err := storage.GetUserByID(1, ctx.DbMap)
	checkTestErr(err)

	salt = mUser.Salt
	assert.True(t, salt != "")
}

func TestInitUser(t *testing.T) {
	ctx := initTest(false)

	user, err := InitUser("zhangyc@fenbi.com", ctx)
	checkTestErr(err)
	assert.Equal(t, 2, user.ID)

	user, err = InitUser("zhangyc@fenbi.com", ctx)
	checkTestErr(err)
	assert.Equal(t, 2, user.ID)
}

func TestGetUserByEmail(t *testing.T) {
	ctx := initTest(false)

	user, err := GetUserByEmail("admin@fenbi.com", ctx)
	checkTestErr(err)

	assert.Equal(t, 1, user.ID)
}

func TestGetUserBySalt(t *testing.T) {
	ctx := initTest(false)

	user, err := GetUserBySalt(salt, ctx)
	checkTestErr(err)

	assert.Equal(t, 1, user.ID)
}

func TestQueryUser(t *testing.T) {
	ctx := initTest(false)

	users, err := QueryUser("in@fe", ctx)
	checkTestErr(err)

	assert.Len(t, users, 1)
}
