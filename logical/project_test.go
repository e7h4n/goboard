package logical

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuantiku/goboard/web/vo"
)

var admin *vo.User
var user *vo.User

func TestSaveProject(t *testing.T) {
	ctx := initTest(true)

	admin = &vo.User{Email: "admin@fenbi.com"}
	checkTestErr(SaveUser(admin, ctx))

	user = &vo.User{Email: "user@fenbi.com"}
	checkTestErr(SaveUser(user, ctx))

	ctx.UserID = admin.ID
	project := &vo.Project{Name: "XXX"}
	checkTestErr(SaveProject(project, ctx))
	uuid := project.UUID

	assert.Equal(t, 1, project.ID)

	project.Name = "Project"
	project.UUID = "xxx"
	checkTestErr(SaveProject(project, ctx))

	assert.Equal(t, uuid, project.UUID)
	assert.Equal(t, "Project", project.Name)
}

func TestGetProject(t *testing.T) {
	ctx := initTest(false)

	ctx.UserID = admin.ID
	project, err := GetProject(1, ctx)
	checkTestErr(err)
	assert.Equal(t, "Project", project.Name)

	ctx.UserID = user.ID
	_, err = GetProject(1, ctx)
	assert.NotNil(t, err)
}

func TestQueryProject(t *testing.T) {
	ctx := initTest(false)

	ctx.UserID = admin.ID
	projects, err := QueryProject(ctx)
	checkTestErr(err)
	assert.Len(t, projects, 1)

	ctx.UserID = user.ID
	projects, err = QueryProject(ctx)
	checkTestErr(err)
	assert.Len(t, projects, 0)
}

func TestRemoveProject(t *testing.T) {
	ctx := initTest(false)

	ctx.UserID = user.ID
	err := RemoveProject(1, ctx)
	assert.NotNil(t, err)

	ctx.UserID = admin.ID
	err = RemoveProject(1, ctx)
	checkTestErr(err)

	projects, err := QueryProject(ctx)
	checkTestErr(err)
	assert.Len(t, projects, 0)
}
