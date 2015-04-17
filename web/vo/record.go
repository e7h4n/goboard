package vo

import (
	"fmt"
	"time"

	"github.com/yuantiku/goboard/storage"
)

// Record is view object for storage.Record
type Record struct {
	ID           int64     `json:"id"`
	DataSourceID int       `json:"dataSourceId"`
	Value        float64   `json:"value"`
	Year         int       `json:"year"`
	Month        int       `json:"month"`
	Day          int       `json:"day"`
	Hour         int       `json:"hour"`
	Minute       int       `json:"minute"`
	Second       int       `json:"second"`
	DateTime     time.Time `json:"dateTime"`
	Dim1         string
	Dim2         string
	Dim3         string
}

// Model convert vo to storage model
func (p *Record) Model() (record *storage.Record) {
	return &storage.Record{
		ID:           p.ID,
		DataSourceID: p.DataSourceID,
		Value:        p.Value,
		Year:         p.Year,
		Month:        p.Month,
		Day:          p.Day,
		Hour:         p.Hour,
		Minute:       p.Minute,
		Second:       p.Second,
		DateTime:     p.DateTime,
		Dim1:         p.Dim1,
		Dim2:         p.Dim2,
		Dim3:         p.Dim3}
}

// NewRecord convert storage model to vo
func NewRecord(r *storage.Record) (record *Record) {
	return &Record{
		ID:           r.ID,
		DataSourceID: r.DataSourceID,
		Value:        r.Value,
		Year:         r.Year,
		Month:        r.Month,
		Day:          r.Day,
		Hour:         r.Hour,
		Minute:       r.Minute,
		Second:       r.Second,
		DateTime:     r.DateTime,
		Dim1:         r.Dim1,
		Dim2:         r.Dim2,
		Dim3:         r.Dim3}
}

// ConvertDimensionFromKey convert a string-object map from dim1.Key to Dim1 by dimension configs
func ConvertDimensionFromKey(dimensions []DimensionConfig, jsonMap map[string]interface{}) {
	if len(dimensions) > 0 {
		for i, v := range dimensions {
			dimKey := fmt.Sprintf("Dim%d", i+1)
			if val, ok := jsonMap[v.Key]; ok {
				jsonMap[dimKey] = val
				delete(jsonMap, v.Key)
			}
		}
	}
}

// ConvertDimensionToKey convert a string-object map from Dim1 to dim1.Key by dimension configs
func ConvertDimensionToKey(dimensions []DimensionConfig, jsonMap map[string]interface{}) {
	if len(dimensions) > 0 {
		for i, v := range dimensions {
			dimKey := fmt.Sprintf("Dim%d", i+1)
			if val, ok := jsonMap[dimKey]; ok {
				jsonMap[v.Key] = val
				delete(jsonMap, dimKey)
			}
		}
	}
}
