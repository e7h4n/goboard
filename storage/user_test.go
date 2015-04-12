package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveUser(t *testing.T) {
	dbmap := InitTestDB(true)
	user := &User{Email: "zhangyc@fenbi.com", Salt: "AAABBB"}
	err := user.Save(dbmap)
	checkTestErr(err)

	assert.Equal(t, 1, user.ID)
}

func TestGetUserByID(t *testing.T) {
	dbmap := InitTestDB(false)

	user := GetUserByID(1, dbmap)

	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
}

func TestGetAllUser(t *testing.T) {
	dbmap := InitTestDB(false)

	users, err := GetAllUser(dbmap)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestQueryUserByEmail(t *testing.T) {
	dbmap := InitTestDB(false)

	user := &User{Email: "gcz@fenbi.com", Salt: "AAABBBDD"}
	err := user.Save(dbmap)
	checkTestErr(err)
	assert.Equal(t, 2, user.ID)

	users, err := QueryUserByEmail("zhangyc", dbmap)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "zhangyc@fenbi.com", users[0].Email)
}

func TestUpdateUser(t *testing.T) {
	dbmap := InitTestDB(false)

	user := GetUserByID(1, dbmap)

	user.Email = "xuhf@fenbi.com"
	err := user.Save(dbmap)
	checkTestErr(err)

	user = GetUserByID(1, dbmap)

	assert.Equal(t, "xuhf@fenbi.com", user.Email)
}

func TestInitUser(t *testing.T) {
	dbmap := InitTestDB(false)

	user, err := InitUser("shenly@fenbi.com", dbmap)
	assert.Nil(t, err)
	assert.Equal(t, 3, user.ID)

	user, err = InitUser("shenly@fenbi.com", dbmap)
	assert.Nil(t, err)
	assert.Equal(t, 3, user.ID)

	user, err = InitUser("xuhf@fenbi.com", dbmap)
	assert.Nil(t, err)
	assert.Equal(t, 1, user.ID)
}
