package logical

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/perfectworks/goboard/web/vo"
)

func TestSaveFolder(t *testing.T) {
	ctx := initTest(true)

	folder := &vo.Folder{Name: "root"}
	err := SaveFolder(1, vo.FolderRoot, folder, ctx)
	checkTestErr(err)
	assert.Equal(t, 1, folder.ID)

	folder = &vo.Folder{Name: "root2"}
	err = SaveFolder(1, vo.FolderRoot, folder, ctx)
	checkTestErr(err)
	assert.Equal(t, 2, folder.ID)

	folder.Name = "root2_"
	err = SaveFolder(1, vo.FolderRoot, folder, ctx)
	checkTestErr(err)
	assert.Equal(t, "root2_", folder.Name)

	folder = &vo.Folder{Name: "sub"}
	err = SaveFolder(1, 2, folder, ctx)
	checkTestErr(err)
	assert.Equal(t, 3, folder.ID)
}

func TestGetFolder(t *testing.T) {
	ctx := initTest(false)

	folder, err := GetFolder(1, ctx)
	checkTestErr(err)
	assert.Equal(t, "root", folder.Name)
}

func TestQueryFolder(t *testing.T) {
	ctx := initTest(false)

	folders, err := QueryFolder(1, vo.FolderRoot, ctx)
	checkTestErr(err)
	assert.Len(t, folders, 2)

	folders, err = QueryFolder(1, 2, ctx)
	checkTestErr(err)
	assert.Len(t, folders, 1)
}

func TestRemoveFolder(t *testing.T) {
	ctx := initTest(false)

	err := RemoveFolder(1, ctx)
	checkTestErr(err)

	folders, err := QueryFolder(1, vo.FolderRoot, ctx)
	checkTestErr(err)
	assert.Len(t, folders, 1)
}
