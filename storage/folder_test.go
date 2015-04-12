package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveFolder(t *testing.T) {
	dbmap := InitTestDB(true)

	project := &Project{Name: "Proj"}
	checkTestErr(project.Save(dbmap))

	folder := &Folder{Name: "RootFolderA", ProjectID: project.ID}
	checkTestErr(folder.Save(dbmap))

	assert.Equal(t, 1, folder.ID)

	folder = &Folder{Name: "RootFolderB", ProjectID: project.ID}
	checkTestErr(folder.Save(dbmap))

	assert.Equal(t, 2, folder.ID)

	folder = &Folder{Name: "subfolder", ParentID: folder.ID, ProjectID: project.ID}
	checkTestErr(folder.Save(dbmap))

	assert.Equal(t, 3, folder.ID)
}

func TestQueryFolder(t *testing.T) {
	dbmap := InitTestDB(false)

	rootFolders, err := QueryFolder(1, FolderRoot, dbmap)
	checkTestErr(err)
	assert.Len(t, rootFolders, 2)

	subFolders, err := QueryFolder(1, rootFolders[1].ID, dbmap)
	checkTestErr(err)
	assert.Len(t, subFolders, 1)
}
