package vo

import (
	"encoding/json"
	"time"

	"github.com/yuantiku/goboard/storage"
)

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

type DataSourceConfig struct {
	Dimensions []DimensionConfig `json:"dimensions"`
}

type DimensionConfig struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Type string `json:"type"`
}

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
