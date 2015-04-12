package vo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDashboardModel(t *testing.T) {
	grids := make([]DashboardLayout, 2)
	grids[0] = DashboardLayout{ID: 1, FirstGrid: LayoutGrid{1, 2}, LastGrid: LayoutGrid{3, 4}}
	grids[1] = DashboardLayout{ID: 2, FirstGrid: LayoutGrid{3, 4}, LastGrid: LayoutGrid{5, 6}}

	config := DashboardConfig{Layout: grids}
	w := &Dashboard{Config: config, CreatedAt: time.Now(), UpdatedAt: time.Now(), Private: false, ID: 1, ProjectID: 1}
	dashboard, err := w.Model()
	checkTestErr(err)

	assert.NotNil(t, dashboard)
	assert.Equal(t, w.ID, dashboard.ID)
	assert.Equal(t, w.ProjectID, dashboard.ProjectID)
	assert.Equal(t, w.CreatedAt, dashboard.CreatedAt)
	assert.Equal(t, w.UpdatedAt, dashboard.UpdatedAt)
	assert.Equal(t, w.Private, dashboard.Private)
	assert.Equal(t, `{"layout":[{"id":1,"firstGrid":[1,2],"lastGrid":[3,4]},{"id":2,"firstGrid":[3,4],"lastGrid":[5,6]}]}`, dashboard.Config)
}
