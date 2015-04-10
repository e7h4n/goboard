package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSaveQuery(t *testing.T) {
	dbmap := initTestDb(true)

	project := &Project{Name: "Demo Project"}
	checkTestErr(project.Save(dbmap))

	dataSource := &DataSource{Name: "DS", ProjectID: project.ID, Key: "test.key", Increment: false}
	checkTestErr(dataSource.Save(dbmap))

	record := &Record{DataSourceID: dataSource.ID, Year: 2012, Month: 5, Day: 10, Hour: 10, Value: 100.0}
	checkTestErr(record.Save(dbmap))

	assert.Equal(t, 1, record.ID)
	assert.Equal(t, 100.0, record.Value)
	assert.True(t, record.CreatedAt.Unix() > 0)
	assert.True(t, record.UpdatedAt.Unix() > 0)
	assert.True(t, record.DateTime.Unix() > 0)

	record = &Record{DataSourceID: dataSource.ID, Year: record.Year, Month: record.Month, Day: record.Day, Hour: record.Hour, Value: 99.0}
	checkTestErr(record.Save(dbmap))
	assert.Equal(t, 1, record.ID)

	record = &Record{DataSourceID: dataSource.ID, Year: record.Year, Month: record.Month, Day: record.Day, Hour: record.Hour, Second: 1, Value: 99.0, Dim1: "A", Dim2: "B", Dim3: "C"}
	checkTestErr(record.Save(dbmap))
	assert.Equal(t, 2, record.ID)
}

func TestQueryRecord(t *testing.T) {
	dbmap := initTestDb(false)
	localLoc, err := time.LoadLocation("Local")
	checkTestErr(err)

	startTime := time.Date(2012, 5, 10, 0, 0, 0, 0, localLoc)
	endTime := time.Date(2012, 5, 11, 0, 0, 0, 0, localLoc)

	filter := RecordFilter{StartTime: startTime, EndTime: endTime, TimeLimit: 2, Dim1: RecordDimFilter("A"), Dim2: RecordDimFilter("B"), Dim3: RecordDimFilter("C")}
	records, err := QueryRecord(1, filter, dbmap)
	checkTestErr(err)

	assert.NotNil(t, records)
	assert.Len(t, records, 1)

	filter = RecordFilter{StartTime: startTime, EndTime: endTime}
	records, err = QueryRecord(1, filter, dbmap)
	checkTestErr(err)

	assert.Len(t, records, 2)

	filter = RecordFilter{StartTime: startTime, EndTime: endTime, Count: 1}
	records, err = QueryRecord(1, filter, dbmap)
	checkTestErr(err)

	assert.Len(t, records, 1)
	assert.Equal(t, 2, records[0].ID)

	filter = RecordFilter{StartTime: startTime, EndTime: endTime, Count: 1, Offset: 1, OrderBy: "id"}
	records, err = QueryRecord(1, filter, dbmap)
	checkTestErr(err)

	assert.Len(t, records, 1)
	assert.Equal(t, 1, records[0].ID)
}

func TestGetRecord(t *testing.T) {
	dbmap := initTestDb(false)

	record, err := GetRecord(1, dbmap)
	checkTestErr(err)

	assert.NotNil(t, record)
}

func TestRemoveRecord(t *testing.T) {
	dbmap := initTestDb(false)

	record, err := GetRecord(1, dbmap)
	checkTestErr(err)

	err = record.Remove(dbmap)
	checkTestErr(err)

	records, err := QueryRecord(1, RecordFilter{}, dbmap)
	checkTestErr(err)

	assert.Len(t, records, 1)
}
