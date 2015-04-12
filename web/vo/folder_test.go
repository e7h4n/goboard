package vo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFolderModel(t *testing.T) {
	f := &Folder{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Name",
		ParentID:  1,
		ProjectID: 1}
	folder := f.Model()

	assert.NotNil(t, folder)
	assert.Equal(t, f.ID, folder.ID)
	assert.Equal(t, f.Name, folder.Name)
	assert.Equal(t, f.CreatedAt, folder.CreatedAt)
	assert.Equal(t, f.UpdatedAt, folder.UpdatedAt)
	assert.Equal(t, f.ProjectID, folder.ProjectID)
	assert.Equal(t, f.ParentID, folder.ParentID)

	f = NewFolder(folder)
	assert.Equal(t, f.ID, folder.ID)
	assert.Equal(t, f.Name, folder.Name)
	assert.Equal(t, f.CreatedAt, folder.CreatedAt)
	assert.Equal(t, f.UpdatedAt, folder.UpdatedAt)
	assert.Equal(t, f.ProjectID, folder.ProjectID)
	assert.Equal(t, f.ParentID, folder.ParentID)
}
