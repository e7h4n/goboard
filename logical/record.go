package logical

import (
	"encoding/json"
	"fmt"
	"time"

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
	jsonMap := make(map[string]interface{}, 16)

	jsonMap["id"] = record.ID
	jsonMap["dataSourceId"] = record.DataSourceID
	jsonMap["value"] = record.Value
	jsonMap["year"] = record.Year
	jsonMap["month"] = record.Month
	jsonMap["day"] = record.Day
	jsonMap["hour"] = record.Hour
	jsonMap["minute"] = record.Minute
	jsonMap["second"] = record.Second
	jsonMap["dateTime"] = record.DateTime
	jsonMap["dim1"] = record.Dim1
	jsonMap["dim2"] = record.Dim2
	jsonMap["dim3"] = record.Dim3

	dataSource, err := GetDataSource(record.DataSourceID, ctx)
	if err != nil {
		return "", err
	}

	if dataSource.Config.Dimensions != nil && len(dataSource.Config.Dimensions) > 0 {
		for i, v := range dataSource.Config.Dimensions {
			dimKey := fmt.Sprintf("dim%d", i+1)
			if val, ok := jsonMap[dimKey]; ok {
				jsonMap[v.Key] = val
				delete(jsonMap, dimKey)
			}
		}
	}

	bRecord, err := json.Marshal(jsonMap)
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

	record = &vo.Record{}

	record.ID = int64(jsonMap["id"].(float64))
	record.DataSourceID = int(jsonMap["dataSourceId"].(float64))
	record.Value = jsonMap["value"].(float64)
	record.Year = int(jsonMap["year"].(float64))
	record.Month = int(jsonMap["month"].(float64))
	record.Day = int(jsonMap["day"].(float64))
	record.Hour = int(jsonMap["hour"].(float64))
	record.Minute = int(jsonMap["minute"].(float64))
	record.Second = int(jsonMap["second"].(float64))
	sTime := jsonMap["dateTime"].(string)
	record.DateTime, err = time.Parse(time.RFC3339, sTime)
	if err != nil {
		return nil, err
	}

	dataSource, err := GetDataSource(record.DataSourceID, ctx)
	if err != nil {
		return nil, err
	}

	if dataSource.Config.Dimensions != nil && len(dataSource.Config.Dimensions) > 0 {
		for i, v := range dataSource.Config.Dimensions {
			if val, ok := jsonMap[v.Key]; ok {
				if i == 0 {
					record.Dim1 = val.(string)
				} else if i == 1 {
					record.Dim2 = val.(string)
				} else if i == 2 {
					record.Dim3 = val.(string)
				}
			}
		}
	}

	return
}
