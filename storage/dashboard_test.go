package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateDashboard(t *testing.T) {
	dbmap := initTestDb(true)

	user := &User{Email: "foo@fenbi.com"}
	err := user.Save(dbmap)
	checkTestErr(err)

	project := &Project{Name: "FooProj"}
	checkTestErr(project.Save(dbmap))

	dashboard := &Dashboard{Name: "FooBoard", ProjectID: project.ID, Private: false, OwnerID: user.ID}
	checkTestErr(dashboard.Save(dbmap))

	assert.Equal(t, 1, dashboard.ID)

	dashboard = GetDashboardByID(1, dbmap)

	assert.Equal(t, "FooBoard", dashboard.Name)
	assert.True(t, dashboard.CreatedAt.Unix() > 0)
	assert.True(t, dashboard.UpdatedAt.Unix() > 0)

	updatedAt := dashboard.UpdatedAt
	dashboard.Name = "BarBoard"
	checkTestErr(dashboard.Save(dbmap))

	dashboard = GetDashboardByID(1, dbmap)

	assert.True(t, dashboard.UpdatedAt.After(updatedAt))
}

func TestQueryDashboardByUser(t *testing.T) {
	dbmap := initTestDb(false)

	users, err := GetAllUser(dbmap)
	checkTestErr(err)

	user := users[0]

	projects, err := GetAllProject(dbmap)
	checkTestErr(err)

	project := projects[0]

	dashboards, err := QueryDashboardByUser(project.ID, user.ID, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 1)

	privateUser := &User{Email: "bar@fenbi.com"}
	checkTestErr(privateUser.Save(dbmap))

	privateBoard := &Dashboard{Name: "PrivateBoard", OwnerID: privateUser.ID, ProjectID: project.ID, Private: true}
	checkTestErr(privateBoard.Save(dbmap))

	dashboards, err = QueryDashboardByUser(project.ID, user.ID, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 1)

	dashboards, err = QueryDashboardByUser(project.ID, privateUser.ID, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 2)
}

func TestRemoveDashboard(t *testing.T) {
	dbmap := initTestDb(false)

	dashboards, err := QueryDashboardByUser(1, 2, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 2)

	checkTestErr(dashboards[0].Remove(dbmap))

	dashboards, err = QueryDashboardByUser(1, 2, dbmap)
	checkTestErr(err)

	assert.Len(t, dashboards, 1)
}
