package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveRole(t *testing.T) {
	dbmap := InitTestDB(true)

	role := &Role{Name: "test", Scope: RoleGlobal}
	err := role.Save(dbmap)
	checkTestErr(err)

	assert.Equal(t, 3, role.ID)
}
