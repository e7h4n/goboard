package vo

import (
	"encoding/json"
	"time"

	"github.com/perfectworks/goboard/storage"
)

// Dashboard is view object for storage.Dashboard
type Dashboard struct {
	ID        int             `json:"id"`
	OwnerID   int             `json:"ownerId"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Name      string          `json:"name"`
	Config    DashboardConfig `json:"config"`
	ProjectID int             `json:"projectId"`
	Private   bool            `json:"provate"`
}

// DashboardConfig is config for dashboard
type DashboardConfig struct {
	Layout []DashboardLayout `json:"layout"`
}

// DashboardLayout is layout config for dashboard
type DashboardLayout struct {
	ID        int        `json:"id"`
	FirstGrid LayoutGrid `json:"firstGrid"`
	LastGrid  LayoutGrid `json:"lastGrid"`
}

// LayoutGrid is layout grid infomations for dashboard
type LayoutGrid [2]int

// Model convert vo to storage model
func (d *Dashboard) Model() (dashboard *storage.Dashboard, err error) {
	dashboard = &storage.Dashboard{
		ID:        d.ID,
		OwnerID:   d.OwnerID,
		Name:      d.Name,
		Private:   d.Private,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ProjectID: d.ProjectID}

	config, err := json.Marshal(d.Config)
	if err != nil {
		return
	}

	dashboard.Config = string(config)
	return
}

// NewDashboard convert storage model to vo
func NewDashboard(d *storage.Dashboard) (dashboard *Dashboard, err error) {
	dashboard = &Dashboard{
		ID:        d.ID,
		OwnerID:   d.OwnerID,
		Name:      d.Name,
		Private:   d.Private,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		ProjectID: d.ProjectID}

	err = json.Unmarshal([]byte(d.Config), &dashboard.Config)
	return
}
