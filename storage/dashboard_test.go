package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateDashboard(t *testing.T) {
	dbmap := InitTestDB(true)

	user := &User{Email: "foo@fenbi.com"}
	err := user.Save(dbmap)
	checkTestErr(err)

	project := &Project{Name: "FooProj"}
	checkTestErr(project.Save(dbmap))

	dashboard := &Dashboard{Name: "FooBoard", ProjectID: project.ID, Private: false, OwnerID: user.ID}
	checkTestErr(dashboard.Save(dbmap))

	assert.Equal(t, 1, dashboard.ID)

	dashboard, err = GetDashboard(1, user.ID, dbmap)
	checkTestErr(err)

	assert.Equal(t, "FooBoard", dashboard.Name)
	assert.True(t, dashboard.CreatedAt.Unix() > 0)
	assert.True(t, dashboard.UpdatedAt.Unix() > 0)

	updatedAt := dashboard.UpdatedAt
	dashboard.Name = "BarBoard"
	checkTestErr(dashboard.Save(dbmap))

	dashboard, err = GetDashboard(1, user.ID, dbmap)
	checkTestErr(err)

	assert.True(t, dashboard.UpdatedAt.After(updatedAt))
}

func TestQueryDashboardByUser(t *testing.T) {
	dbmap := InitTestDB(false)

	users, err := GetAllUser(dbmap)
	checkTestErr(err)

	user := users[0]

	projects, err := QueryProject(user.ID, dbmap)
	checkTestErr(err)

	assert.Len(t, projects, 1)
	project := projects[0]

	dashboards, err := QueryDashboard(project.ID, user.ID, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 1)

	privateUser := &User{Email: "bar@fenbi.com"}
	checkTestErr(privateUser.Save(dbmap))

	privateBoard := &Dashboard{Name: "PrivateBoard", OwnerID: privateUser.ID, ProjectID: project.ID, Private: true}
	checkTestErr(privateBoard.Save(dbmap))

	dashboards, err = QueryDashboard(project.ID, user.ID, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 1)

	dashboards, err = QueryDashboard(project.ID, privateUser.ID, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 2)
}

func TestRemoveDashboard(t *testing.T) {
	dbmap := InitTestDB(false)

	dashboards, err := QueryDashboard(1, 2, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 2)

	checkTestErr(dashboards[0].Remove(dbmap))

	dashboards, err = QueryDashboard(1, 2, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 1)
}
