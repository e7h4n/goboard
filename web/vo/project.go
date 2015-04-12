package vo

import (
	"time"

	"github.com/yuantiku/goboard/storage"
)

type Project struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (p *Project) Model() (project *storage.Project, err error) {
	project = &storage.Project{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		UUID:      p.UUID,
		Name:      p.Name}

	return
}
