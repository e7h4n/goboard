package models

import (
	"errors"
	"time"

	"gopkg.in/gorp.v1"
)

const (
	// FolderRoot is an id of root folder
	FolderRoot = 0
)

// Folder makes a folder structure of DataSource
type Folder struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string    `db:"name"`
	ParentID  int       `db:"parent_id"`
	ProjectID int       `db:"project_id"`
}

func initFolderTable(dbmap *gorp.DbMap) {
	folderTable := dbmap.AddTableWithName(Folder{}, "folders")
	folderTable.SetKeys(true, "id")
	folderTable.ColMap("name").SetUnique(true)
}

// GetFolderByID will retrieve folder which specified by id
func GetFolderByID(folderID int, dbmap *gorp.DbMap) (folder *Folder) {
	err := dbmap.SelectOne(&folder, "select * from folder where id = ?", folderID)
	if err != nil {
		return nil
	}

	return
}

// QueryFolder will retrieve folders by project and parent folder
func QueryFolder(projectID int, parentID int, dbmap *gorp.DbMap) (folders []Folder, err error) {
	_, err = dbmap.Select(&folders, "select * from folders where project_id = ? and parent_id = ?", projectID, parentID)
	return
}

// Save will insert or update database record
func (f *Folder) Save(dbmap *gorp.DbMap) (err error) {
	f.UpdatedAt = time.Now()

	var affectCount int64
	if f.ID > 0 {
		affectCount, err = dbmap.Update(f)
	} else {
		f.CreatedAt = f.UpdatedAt
		err = dbmap.Insert(f)
		affectCount = 1
	}

	if err != nil {
		return err
	}

	if affectCount == 0 {
		return errors.New("failed to save folder, affectCount = 0")
	}

	return
}

// Remove will remove database record.
func (f *Folder) Remove(dbmap *gorp.DbMap) (err error) {
	_, err = dbmap.Delete(f)
	return
}
