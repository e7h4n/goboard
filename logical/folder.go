package logical

import (
	"github.com/perfectworks/goboard/storage"
	"github.com/perfectworks/goboard/web/vo"
)

// QueryFolder will retrieve folders by project and parent
func QueryFolder(projectID int, parentID int, ctx *Context) (folders []vo.Folder, err error) {
	mFolders, err := storage.QueryFolder(projectID, parentID, ctx.DbMap)

	folders = make([]vo.Folder, len(mFolders))
	for i, v := range mFolders {
		folder := vo.NewFolder(&v)
		folders[i] = *folder
	}

	return
}

// GetFolder will retrieve folder by id
func GetFolder(folderID int, ctx *Context) (folder *vo.Folder, err error) {
	mFolder, err := storage.GetFolder(folderID, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	folder = vo.NewFolder(mFolder)

	return
}

// SaveFolder will update or create a folder
func SaveFolder(projectID int, parentID int, folder *vo.Folder, ctx *Context) (err error) {
	if folder.ID > 0 {
		mFolder, err := storage.GetFolder(folder.ID, ctx.DbMap)
		if err != nil {
			return err
		}

		folder.ProjectID = mFolder.ProjectID
	} else {
		folder.ProjectID = projectID
	}

	folder.ParentID = parentID

	mFolder := folder.Model()

	err = mFolder.Save(ctx.DbMap)
	if err != nil {
		return err
	}

	newFolder := vo.NewFolder(mFolder)

	*folder = *newFolder

	return
}

// RemoveFolder delete a folder
func RemoveFolder(folderID int, ctx *Context) (err error) {
	folder := &storage.Folder{ID: folderID}
	err = folder.Remove(ctx.DbMap)
	return
}
