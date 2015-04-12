package vo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProjectModel(t *testing.T) {
	p := &Project{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Name",
		UUID:      "AAA"}
	project := p.Model()

	assert.NotNil(t, project)
	assert.Equal(t, p.ID, project.ID)
	assert.Equal(t, p.Name, project.Name)
	assert.Equal(t, p.CreatedAt, project.CreatedAt)
	assert.Equal(t, p.UpdatedAt, project.UpdatedAt)
	assert.Equal(t, p.UUID, project.UUID)

	p = NewProject(project)
	assert.Equal(t, p.ID, project.ID)
	assert.Equal(t, p.Name, project.Name)
	assert.Equal(t, p.CreatedAt, project.CreatedAt)
	assert.Equal(t, p.UpdatedAt, project.UpdatedAt)
	assert.Equal(t, p.UUID, project.UUID)
}
