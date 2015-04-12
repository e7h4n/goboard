package vo

import (
	"time"

	"github.com/yuantiku/goboard/storage"
)

// Project is view object for storage.Project
type Project struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Model convert vo to storage model
func (p *Project) Model() (project *storage.Project) {
	return &storage.Project{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		UUID:      p.UUID,
		Name:      p.Name}
}

// NewProject convert storage model to vo
func NewProject(p *storage.Project) (project *Project) {
	return &Project{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		UUID:      p.UUID,
		Name:      p.Name}
}
