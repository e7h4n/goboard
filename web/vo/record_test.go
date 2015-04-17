package vo

import (
	"testing"
	"time"

	"github.com/fatih/structs"
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
		DateTime:     time.Now(),
		Dim1:         "foo",
		Dim2:         "bar",
		Dim3:         "baz"}
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
	assert.Equal(t, r.Dim1, record.Dim1)
	assert.Equal(t, r.Dim2, record.Dim2)
	assert.Equal(t, r.Dim3, record.Dim3)

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
	assert.Equal(t, r.Dim1, record.Dim1)
	assert.Equal(t, r.Dim2, record.Dim2)
	assert.Equal(t, r.Dim3, record.Dim3)
}

func TestConvertDimensionToKey(t *testing.T) {
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
		DateTime:     time.Now(),
		Dim1:         "foo",
		Dim2:         "bar",
		Dim3:         "baz"}

	m := structs.Map(r)
	dimensions := make([]DimensionConfig, 3)
	dimensions[0] = DimensionConfig{Key: "hello"}
	dimensions[1] = DimensionConfig{Key: "world"}
	dimensions[2] = DimensionConfig{Key: "!"}
	ConvertDimensionToKey(dimensions, m)

	assert.Equal(t, "foo", m["hello"])
	assert.Equal(t, "bar", m["world"])
	assert.Equal(t, "baz", m["!"])

	_, ok := m["Dim1"]
	assert.False(t, ok)
	_, ok = m["Dim2"]
	assert.False(t, ok)
	_, ok = m["Dim3"]
	assert.False(t, ok)
}

func TestConvertDimensionFromKey(t *testing.T) {
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
		DateTime:     time.Now(),
		Dim1:         "foo",
		Dim2:         "bar",
		Dim3:         "baz"}

	m := structs.Map(r)
	dimensions := make([]DimensionConfig, 3)
	dimensions[0] = DimensionConfig{Key: "hello"}
	dimensions[1] = DimensionConfig{Key: "world"}
	dimensions[2] = DimensionConfig{Key: "!"}
	ConvertDimensionToKey(dimensions, m)
	ConvertDimensionFromKey(dimensions, m)

	assert.Equal(t, "foo", m["Dim1"])
	assert.Equal(t, "bar", m["Dim2"])
	assert.Equal(t, "baz", m["Dim3"])

	_, ok := m["hello"]
	assert.False(t, ok)
	_, ok = m["world"]
	assert.False(t, ok)
	_, ok = m["!"]
	assert.False(t, ok)
}
