package logical

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/perfectworks/goboard/web/vo"
)

func TestSaveDataSource(t *testing.T) {
	ctx := initTest(true)

	dataSource := &vo.DataSource{Name: "Foo", Key: "Foo"}
	err := SaveDataSource(1, dataSource, ctx)
	checkTestErr(err)

	assert.Equal(t, 1, dataSource.ID)

	dataSource.Name = "Bar"
	err = SaveDataSource(1, dataSource, ctx)
	checkTestErr(err)

	assert.Equal(t, 1, dataSource.ID)
	assert.Equal(t, "Bar", dataSource.Name)

	dataSource = &vo.DataSource{Name: "Sub1", FolderID: 1, Key: "sub1"}
	checkTestErr(SaveDataSource(1, dataSource, ctx))

	dataSource = &vo.DataSource{Name: "Sub2", FolderID: 1, Key: "sub2"}
	checkTestErr(SaveDataSource(1, dataSource, ctx))
}

func TestGetDataSource(t *testing.T) {
	ctx := initTest(false)

	dataSource, err := GetDataSource(1, ctx)
	checkTestErr(err)

	assert.Equal(t, "Bar", dataSource.Name)
}

func TestQueryDataSource(t *testing.T) {
	ctx := initTest(false)

	dataSources, err := QueryDataSource(1, vo.FolderRoot, ctx)
	checkTestErr(err)
	assert.Len(t, dataSources, 1)

	dataSources, err = QueryDataSource(1, 1, ctx)
	checkTestErr(err)
	assert.Len(t, dataSources, 2)
}

func TestRemoveDataSource(t *testing.T) {
	ctx := initTest(false)

	err := RemoveDataSource(1, ctx)
	checkTestErr(err)

	dataSources, err := QueryDataSource(1, vo.FolderRoot, ctx)
	checkTestErr(err)
	assert.Len(t, dataSources, 0)
}
