package vo

import (
	"time"

	"github.com/yuantiku/goboard/storage"
)

type Folder struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ParentID  int       `json:"parentId"`
	ProjectID int       `json:"projectId"`
}

func (f *Folder) Model() (folder *storage.Folder, err error) {
	folder = &storage.Folder{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		ParentID:  f.ParentID,
		ProjectID: f.ProjectID}

	return
}
