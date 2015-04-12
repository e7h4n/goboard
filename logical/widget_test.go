package logical

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuantiku/goboard/web/vo"
)

func TestSaveWidget(t *testing.T) {
	ctx := initTest(true)

	widget := &vo.Widget{}

	err := SaveWidget(1, widget, ctx)
	checkTestErr(err)

	assert.Equal(t, 1, widget.ID)
}

func TestGetWidget(t *testing.T) {
	ctx := initTest(false)

	widget, err := GetWidget(1, ctx)
	checkTestErr(err)

	assert.Equal(t, 1, widget.DashboardID)
}

func TestQueryWidget(t *testing.T) {
	ctx := initTest(false)

	widgets, err := QueryWidget(1, ctx)
	checkTestErr(err)

	assert.Len(t, widgets, 1)

	widgets, err = QueryWidget(2, ctx)
	checkTestErr(err)

	assert.Len(t, widgets, 0)
}

func TestRemoveWidget(t *testing.T) {
	ctx := initTest(false)

	err := RemoveWidget(1, ctx)
	checkTestErr(err)

	widgets, err := QueryWidget(1, ctx)
	checkTestErr(err)

	assert.Len(t, widgets, 0)
}
