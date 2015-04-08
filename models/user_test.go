package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveUser(t *testing.T) {
	dbmap := initDb(true)
	user := &User{Email: "zhangyc@fenbi.com", Salt: "AAABBB"}
	user.Save(dbmap)

	assert.Equal(t, 1, user.Id)
}

func TestGetUserById(t *testing.T) {
	dbmap := initDb(false)

	user := GetUserById(1, dbmap)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.Id)
}

func TestGetAllUser(t *testing.T) {
	dbmap := initDb(false)

	users := GetAllUser(dbmap)
	assert.Equal(t, 1, len(users))
}

func TestQueryUserByEmail(t *testing.T) {
	dbmap := initDb(false)

	user := &User{Email: "gcz@fenbi.com", Salt: "AAABBBDD"}
	user.Save(dbmap)
	assert.Equal(t, 2, user.Id)

	users := QueryUserByEmail("zhangyc", dbmap)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "zhangyc@fenbi.com", users[0].Email)
}

func TestUpdateUser(t *testing.T) {
	dbmap := initDb(false)

	user := GetUserById(1, dbmap)
	user.Email = "xuhf@fenbi.com"
	user.Update(dbmap)

	user = GetUserById(1, dbmap)

	assert.Equal(t, "xuhf@fenbi.com", user.Email)
}

func TestInitUser(t *testing.T) {
	dbmap := initDb(false)

	user := InitUser("shenly@fenbi.com", dbmap)
	assert.Equal(t, 3, user.Id)

	user = InitUser("shenly@fenbi.com", dbmap)
	assert.Equal(t, 3, user.Id)

	user = InitUser("xuhf@fenbi.com", dbmap)
	assert.Equal(t, 1, user.Id)
}
