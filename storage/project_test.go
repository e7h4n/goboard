package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveProject(t *testing.T) {
	dbmap := InitTestDB(true)

	admin := &User{Email: "zhangyc@fenbi.com"}
	checkTestErr(admin.Save(dbmap))

	user := &User{Email: "user@fenbi.com"}
	checkTestErr(user.Save(dbmap))

	project := &Project{Name: "Demo Project"}
	err := project.Save(dbmap)
	checkTestErr(err)

	assert.Equal(t, 1, project.ID)
	assert.Len(t, project.UUID, 36)

	project.Name = "Foo"
	projectID := project.ID
	err = project.Save(dbmap)
	checkTestErr(err)

	project, err = GetProject(projectID, 1, dbmap)
	checkTestErr(err)

	projects, err := QueryProject(2, dbmap)
	checkTestErr(err)
	assert.Len(t, projects, 0)

	assert.Equal(t, "Foo", project.Name)
}

func TestDeleteProject(t *testing.T) {
	dbmap := InitTestDB(false)

	projects, err := QueryProject(1, dbmap)
	checkTestErr(err)

	assert.Len(t, projects, 1)

	err = projects[0].Remove(dbmap)
	checkTestErr(err)

	projects, err = QueryProject(1, dbmap)
	checkTestErr(err)

	assert.Len(t, projects, 0)
}
