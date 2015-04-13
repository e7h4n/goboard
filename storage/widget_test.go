package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveWidget(t *testing.T) {
	dbmap := InitTestDB(true)

	user := &User{Email: "foo@fenbi.com"}
	err := user.Save(dbmap)
	checkTestErr(err)

	project := &Project{Name: "FooProj"}
	checkTestErr(project.Save(dbmap))

	dashboard := &Dashboard{Name: "FooBoard", ProjectID: project.ID, Private: false, OwnerID: user.ID}
	checkTestErr(dashboard.Save(dbmap))

	widget := &Widget{Type: WidgetSpline, DashboardID: dashboard.ID}
	checkTestErr(widget.Save(dbmap))
}

func TestQueryWidget(t *testing.T) {
	dbmap := InitTestDB(false)

	widgets, err := QueryWidget(1, dbmap)
	checkTestErr(err)

	assert.Len(t, widgets, 1)
}

func TestRemoveWidget(t *testing.T) {
	dbmap := InitTestDB(false)

	widget, err := GetWidget(1, dbmap)
	checkTestErr(err)
	checkTestErr(widget.Remove(dbmap))

	widgets, err := QueryWidget(1, dbmap)
	checkTestErr(err)

	assert.Len(t, widgets, 0)
}
