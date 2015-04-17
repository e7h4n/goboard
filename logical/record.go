package logical

import (
	"encoding/json"
	"log"

	"github.com/yuantiku/goboard/storage"
	"github.com/yuantiku/goboard/web/vo"
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

	log.Println(string(bRecord))
	record = &vo.Record{}
	err = json.Unmarshal(bRecord, record)
	if err != nil {
		return nil, err
	}

	return
}
