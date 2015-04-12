package vo

import (
	"encoding/json"
	"time"

	"github.com/yuantiku/goboard/storage"
)

type Dashboard struct {
	ID        int             `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Name      string          `json:"name"`
	Config    DashboardConfig `json:"config"`
	ProjectID int             `json:"projectId"`
	Private   bool            `json:"provate"`
}

type DashboardConfig struct {
	Layout []DashboardLayout `json:"layout"`
}

type DashboardLayout struct {
	ID        int        `json:"id"`
	FirstGrid LayoutGrid `json:"firstGrid"`
	LastGrid  LayoutGrid `json:"lastGrid"`
}

type LayoutGrid [2]int

func (d *Dashboard) Model() (dashboard *storage.Dashboard, err error) {
	dashboard = &storage.Dashboard{ID: d.ID, Name: d.Name, Private: d.Private, CreatedAt: d.CreatedAt, UpdatedAt: d.UpdatedAt, ProjectID: d.ProjectID}
	config, err := json.Marshal(d.Config)
	if err != nil {
		return
	}

	dashboard.Config = string(config)
	return
}
