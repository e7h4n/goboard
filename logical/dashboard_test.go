package logical

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/perfectworks/goboard/web/vo"
)

func TestSaveDashboard(t *testing.T) {
	ctx := initTest(true)

	dashboard := &vo.Dashboard{Private: true, Name: "Foo"}

	ctx.UserID = 1
	err := SaveDashboard(1, dashboard, ctx)
	checkTestErr(err)

	assert.Equal(t, 1, dashboard.ID)
	assert.Equal(t, "Foo", dashboard.Name)

	dashboard.Name = "Baz"
	err = SaveDashboard(dashboard.ProjectID, dashboard, ctx)
	checkTestErr(err)

	ctx.UserID = 2
	dashboard = &vo.Dashboard{Private: false, Name: "Bar"}
	err = SaveDashboard(1, dashboard, ctx)
	checkTestErr(err)

	assert.Equal(t, 2, dashboard.ID)
}

func TestGetDashboard(t *testing.T) {
	ctx := initTest(false)

	ctx.UserID = 1
	dashboard, err := GetDashboard(1, ctx)
	checkTestErr(err)

	assert.Equal(t, "Baz", dashboard.Name)
}

func TestQueryDashboard(t *testing.T) {
	ctx := initTest(false)

	ctx.UserID = 1
	dashboards, err := QueryDashboard(1, ctx)
	checkTestErr(err)
	assert.Len(t, dashboards, 2)

	ctx.UserID = 2
	dashboards, err = QueryDashboard(1, ctx)
	checkTestErr(err)
	assert.Len(t, dashboards, 1)
}

func TestRemoveDashboard(t *testing.T) {
	ctx := initTest(false)
	ctx.UserID = 2
	err := RemoveDashboard(2, ctx)
	checkTestErr(err)

	err = RemoveDashboard(1, ctx)
	assert.NotNil(t, err)
}
