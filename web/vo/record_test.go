package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	f := &User{
		ID:    1,
		Email: "zhangyc@fenbi.com",
		Salt:  "AAA"}
	user, err := f.Model()
	checkTestErr(err)

	assert.NotNil(t, user)
	assert.Equal(t, f.ID, user.ID)
	assert.Equal(t, f.Email, user.Email)
	assert.Equal(t, f.Salt, user.Salt)
}
