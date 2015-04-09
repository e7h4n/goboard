package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveDataSource(t *testing.T) {
	dbmap := initTestDb(true)

	project := &Project{Name: "Demo Project"}
	checkTestErr(project.Save(dbmap))

	dataSource := &DataSource{Name: "DS", ProjectID: project.ID, Key: "test.key", Increment: false}
	checkTestErr(dataSource.Save(dbmap))

	assert.Equal(t, 1, dataSource.ID)
	assert.True(t, dataSource.CreatedAt.Unix() > 0)

	assert.Equal(t, FolderRoot, dataSource.FolderID)
}

func TestQueryDataSource(t *testing.T) {
	dbmap := initTestDb(false)

	dataSources, err := QueryDataSource(1, FolderRoot, dbmap)
	checkTestErr(err)

	assert.Len(t, dataSources, 1)
}

func TestGetDataSourceByID(t *testing.T) {
	dbmap := initTestDb(false)

	dataSource := GetDataSourceByID(1, dbmap)

	assert.NotNil(t, dataSource)
}

func TestRemoveDataSource(t *testing.T) {
	dbmap := initTestDb(false)

	dataSource := GetDataSourceByID(1, dbmap)

	checkTestErr(dataSource.Remove(dbmap))

	dataSource = GetDataSourceByID(1, dbmap)

	assert.Nil(t, dataSource)
}
