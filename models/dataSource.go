package models

import (
	"errors"
	"time"

	"gopkg.in/gorp.v1"
)

// A DataSource is a set of records
type DataSource struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string    `db:"name"`
	ProjectID int       `db:"project_id"`
	Key       string    `db:"key"`
	FolderID  int       `db:"folder_id"`
	Config    string    `db:"config"`
	Increment bool      `db:"increment"`
}

func initDataSourceTable(dbmap *gorp.DbMap) {
	dataSourceTable := dbmap.AddTableWithName(DataSource{}, "data_sources")
	dataSourceTable.SetKeys(true, "id")
	dataSourceTable.SetUniqueTogether("key", "project_id")
	dataSourceTable.ColMap("key").SetNotNull(true)
	dataSourceTable.ColMap("name").SetNotNull(true)
}

// QueryDataSource will retrieve DataSources by specified project and folder
func QueryDataSource(projectID int, folderID int, dbmap *gorp.DbMap) (dataSources []DataSource, err error) {
	_, err = dbmap.Select(&dataSources, "select * from data_sources where project_id = ? and folder_id = ?", projectID, folderID)
	return
}

// GetDataSourceByID retrieve specified DataSource by id
func GetDataSourceByID(dataSourceID int, dbmap *gorp.DbMap) (dataSource *DataSource) {
	err := dbmap.SelectOne(&dataSource, "select * from data_sources where id = ?", dataSourceID)
	if err != nil {
		return nil
	}

	return
}

// Save will insert a new data source record or update an existed data source record from database.
func (ds *DataSource) Save(dbmap *gorp.DbMap) (err error) {
	ds.UpdatedAt = time.Now()

	var affectCount int64
	if ds.ID > 0 {
		affectCount, err = dbmap.Update(ds)
	} else {
		ds.CreatedAt = ds.UpdatedAt
		err = dbmap.Insert(ds)
		affectCount = 1
	}

	if err != nil {
		return err
	}

	if affectCount == 0 {
		return errors.New("failed to save data source, affectCount = 0")
	}

	return
}

// Remove will delete the record of DataSource from database
func (ds *DataSource) Remove(dbmap *gorp.DbMap) (err error) {
	_, err = dbmap.Delete(ds)
	return
}
