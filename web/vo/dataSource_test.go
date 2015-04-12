package vo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDataSourceModel(t *testing.T) {
	dimensions := make([]DimensionConfig, 2)
	dimensions[0] = DimensionConfig{Key: "Key", Name: "Name", Type: "Type"}
	dimensions[1] = DimensionConfig{Key: "Key", Name: "Name", Type: "Type"}

	config := DataSourceConfig{Dimensions: dimensions}
	w := &DataSource{
		ID:        1,
		Name:      "Name",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ProjectID: 1,
		Key:       "Key",
		FolderID:  1,
		Increment: true,
		Config:    config}
	dataSource, err := w.Model()
	checkTestErr(err)

	assert.NotNil(t, dataSource)
	assert.Equal(t, w.ID, dataSource.ID)
	assert.Equal(t, w.Name, dataSource.Name)
	assert.Equal(t, w.CreatedAt, dataSource.CreatedAt)
	assert.Equal(t, w.UpdatedAt, dataSource.UpdatedAt)
	assert.Equal(t, w.ProjectID, dataSource.ProjectID)
	assert.Equal(t, w.Key, dataSource.Key)
	assert.Equal(t, w.FolderID, dataSource.FolderID)
	assert.Equal(t, w.Increment, dataSource.Increment)
	assert.Equal(t, `{"dimensions":[{"key":"Key","name":"Name","type":"Type"},{"key":"Key","name":"Name","type":"Type"}]}`, dataSource.Config)

	w, err = NewDataSource(dataSource)
	checkTestErr(err)
	assert.Equal(t, w.ID, dataSource.ID)
	assert.Equal(t, w.Name, dataSource.Name)
	assert.Equal(t, w.CreatedAt, dataSource.CreatedAt)
	assert.Equal(t, w.UpdatedAt, dataSource.UpdatedAt)
	assert.Equal(t, w.ProjectID, dataSource.ProjectID)
	assert.Equal(t, w.Key, dataSource.Key)
	assert.Equal(t, w.FolderID, dataSource.FolderID)
	assert.Equal(t, w.Increment, dataSource.Increment)
}
