package vo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProjectModel(t *testing.T) {
	f := &Project{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Name",
		UUID:      "AAA"}
	project, err := f.Model()
	checkTestErr(err)

	assert.NotNil(t, project)
	assert.Equal(t, f.ID, project.ID)
	assert.Equal(t, f.Name, project.Name)
	assert.Equal(t, f.CreatedAt, project.CreatedAt)
	assert.Equal(t, f.UpdatedAt, project.UpdatedAt)
	assert.Equal(t, f.UUID, project.UUID)
}
