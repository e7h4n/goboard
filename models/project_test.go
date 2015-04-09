package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveProject(t *testing.T) {
	dbmap := initTestDb(true)

	project := &Project{Name: "Demo Project"}
	err := project.Save(dbmap)
	checkTestErr(err)

	assert.Equal(t, 1, project.ID)
	assert.Len(t, project.UUID, 36)

	project.Name = "Foo"
	projectID := project.ID
	err = project.Save(dbmap)
	checkTestErr(err)

	project = GetProjectByID(projectID, dbmap)
	assert.Equal(t, "Foo", project.Name)
}

func TestDeleteProject(t *testing.T) {
	dbmap := initTestDb(false)

	projects, err := GetAllProject(dbmap)
	checkTestErr(err)

	assert.Len(t, projects, 1)

	err = projects[0].Remove(dbmap)
	checkTestErr(err)

	projects, err = GetAllProject(dbmap)
	checkTestErr(err)

	assert.Len(t, projects, 0)
}
