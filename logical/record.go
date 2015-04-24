package logical

import (
	"encoding/json"

	"github.com/perfectworks/goboard/storage"
	"github.com/perfectworks/goboard/web/vo"
)

// SaveRecord create or update a record
func SaveRecord(record *vo.Record, ctx *Context) (err error) {
	mRecord := record.Model()
	err = mRecord.Save(ctx.DbMap)
	if err != nil {
		return err
	}

	*record = *vo.NewRecord(mRecord)
	return
}

// GetRecord retrieve record by id
func GetRecord(recordID int, ctx *Context) (record *vo.Record, err error) {
	mRecord, err := storage.GetRecord(recordID, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	record = vo.NewRecord(mRecord)

	return
}

// QueryRecord retrieve records by dataSourceID and filter
func QueryRecord(dataSourceID int, filter storage.RecordFilter, ctx *Context) (records []vo.Record, err error) {
	mRecords, err := storage.QueryRecord(dataSourceID, filter, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	records = make([]vo.Record, len(mRecords))
	for i, v := range mRecords {
		records[i] = *vo.NewRecord(&v)
	}

	return
}

// ClearRecord remove all records belongs to specified data source
func ClearRecord(dataSourceID int, ctx *Context) (err error) {
	err = storage.ClearRecord(dataSourceID, ctx.DbMap)
	return
}

// RemoveRecord remove a record
func RemoveRecord(recordID int64, ctx *Context) (err error) {
	mRecord := &storage.Record{ID: recordID}
	return mRecord.Remove(ctx.DbMap)
}

// RecordToJSON covert vo.Record to json string
func RecordToJSON(record *vo.Record, ctx *Context) (jRecord string, err error) {
	bRecord, err := json.Marshal(record)
	if err != nil {
		return "", err
	}

	var jsonMap map[string]interface{}
	err = json.Unmarshal(bRecord, &jsonMap)
	if err != nil {
		return "", err
	}

	dataSource, err := GetDataSource(record.DataSourceID, ctx)
	if err != nil {
		return "", err
	}

	vo.ConvertDimensionToKey(dataSource.Config.Dimensions, jsonMap)

	bRecord, err = json.Marshal(jsonMap)
	if err != nil {
		return "", err
	}

	jRecord = string(bRecord)
	return
}

// RecrodFromJSON covert json string to vo.Record
func RecrodFromJSON(jRecord string, ctx *Context) (record *vo.Record, err error) {
	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(jRecord), &jsonMap)
	if err != nil {
		return nil, err
	}

	dataSource, err := GetDataSource(int(jsonMap["dataSourceId"].(float64)), ctx)
	if err != nil {
		return nil, err
	}

	vo.ConvertDimensionFromKey(dataSource.Config.Dimensions, jsonMap)

	bRecord, err := json.Marshal(jsonMap)
	if err != nil {
		return nil, err
	}

	record = &vo.Record{}
	err = json.Unmarshal(bRecord, record)
	if err != nil {
		return nil, err
	}

	return
}
