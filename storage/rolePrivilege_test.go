package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryRolePrivilegesByRoleID(t *testing.T) {
	dbmap := InitTestDB(true)

	user := &User{Email: "zhangyc@fenbi.com"}
	err := user.Save(dbmap)
	checkTestErr(err)

	roles, err := QueryRoleByScope(RoleGlobal, dbmap)
	checkTestErr(err)

	rolePrivileges, err := QueryRolePrivilegesByRoleID(roles[0].ID, dbmap)
	checkTestErr(err)

	assert.Equal(t, 4, len(rolePrivileges))
}
