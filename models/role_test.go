package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllRole(t *testing.T) {
	dbmap := initTestDb(true)

	roles, err := GetAllRole(dbmap)
	checkTestErr(err)

	assert.Equal(t, 2, len(roles))
}
