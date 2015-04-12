package vo

import (
	"time"

	"github.com/yuantiku/goboard/storage"
)

// Folder is view object for storage.Folder
type Folder struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ParentID  int       `json:"parentId"`
	ProjectID int       `json:"projectId"`
}

// Model convert vo to storage model
func (f *Folder) Model() (folder *storage.Folder) {
	return &storage.Folder{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		ParentID:  f.ParentID,
		ProjectID: f.ProjectID}
}

func NewFolder(f *storage.Folder) (folder *Folder) {
	return &Folder{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		ParentID:  f.ParentID,
		ProjectID: f.ProjectID}
}