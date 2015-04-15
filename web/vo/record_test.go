package vo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRecordModel(t *testing.T) {
	r := &Record{
		ID:           1,
		DataSourceID: 1,
		Value:        100,
		Year:         2014,
		Month:        2,
		Day:          1,
		Hour:         1,
		Minute:       10,
		Second:       0,
		DateTime:     time.Now()}
	record := r.Model()

	assert.NotNil(t, record)
	assert.Equal(t, r.ID, record.ID)
	assert.Equal(t, r.DataSourceID, record.DataSourceID)
	assert.Equal(t, r.Value, record.Value)
	assert.Equal(t, r.Year, record.Year)
	assert.Equal(t, r.Month, record.Month)
	assert.Equal(t, r.Day, record.Day)
	assert.Equal(t, r.Hour, record.Hour)
	assert.Equal(t, r.Minute, record.Minute)
	assert.Equal(t, r.Second, record.Second)
	assert.Equal(t, r.DateTime, record.DateTime)

	r = NewRecord(record)
	assert.Equal(t, r.ID, record.ID)
	assert.Equal(t, r.DataSourceID, record.DataSourceID)
	assert.Equal(t, r.Value, record.Value)
	assert.Equal(t, r.Year, record.Year)
	assert.Equal(t, r.Month, record.Month)
	assert.Equal(t, r.Day, record.Day)
	assert.Equal(t, r.Hour, record.Hour)
	assert.Equal(t, r.Minute, record.Minute)
	assert.Equal(t, r.Second, record.Second)
	assert.Equal(t, r.DateTime, record.DateTime)
}
