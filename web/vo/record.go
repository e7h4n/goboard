package vo

import (
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
	Dim1         string    `json:"dim1"`
	Dim2         string    `json:"dim2"`
	Dim3         string    `json:"dim3"`
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
