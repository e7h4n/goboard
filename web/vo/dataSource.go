package vo

import (
	"encoding/json"
	"time"

	"github.com/perfectworks/goboard/storage"
)

// DataSource is view object for storage.DataSource
type DataSource struct {
	ID        int              `json:"id"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Name      string           `json:"name"`
	ProjectID int              `json:"projectId"`
	Key       string           `json:"key"`
	FolderID  int              `json:"folderId"`
	Config    DataSourceConfig `json:"config"`
	Increment bool             `json:"increment"`
}

// DataSourceConfig is data source config
type DataSourceConfig struct {
	Dimensions []DimensionConfig `json:"dimensions"`
}

// DimensionConfig is dimension config for data source
type DimensionConfig struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Model convert vo to storage model
func (ds *DataSource) Model() (dataSource *storage.DataSource, err error) {
	dataSource = &storage.DataSource{
		ID:        ds.ID,
		Name:      ds.Name,
		CreatedAt: ds.CreatedAt,
		UpdatedAt: ds.UpdatedAt,
		ProjectID: ds.ProjectID,
		Key:       ds.Key,
		FolderID:  ds.FolderID,
		Increment: ds.Increment}

	config, err := json.Marshal(ds.Config)
	if err != nil {
		return
	}

	dataSource.Config = string(config)
	return
}

// NewDataSource convert storage model to vo
func NewDataSource(ds *storage.DataSource) (dataSource *DataSource, err error) {
	dataSource = &DataSource{
		ID:        ds.ID,
		Name:      ds.Name,
		CreatedAt: ds.CreatedAt,
		UpdatedAt: ds.UpdatedAt,
		ProjectID: ds.ProjectID,
		Key:       ds.Key,
		FolderID:  ds.FolderID,
		Increment: ds.Increment}

	err = json.Unmarshal([]byte(ds.Config), &dataSource.Config)

	return
}
