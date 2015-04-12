package vo

import (
	"time"

	"github.com/yuantiku/goboard/storage"
)

type Record struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DataSourceID int       `json:"dataSourceId"`
	Value        float64   `json:"value"`
	Year         int       `json:"year"`
	Month        int       `json:"month"`
	Day          int       `json:"day"`
	Hour         int       `json:"hour"`
	Minute       int       `json:"minute"`
	Second       int       `json:"second"`
	DateTime     time.Time `json:"dateTime"`
}

func (p *Record) Model() (record *storage.Record, err error) {
	record = &storage.Record{
		ID:           p.ID,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
		DataSourceID: p.DataSourceID,
		Value:        p.Value,
		Year:         p.Year,
		Month:        p.Month,
		Day:          p.Day,
		Hour:         p.Hour,
		Minute:       p.Minute,
		Second:       p.Second,
		DateTime:     p.DateTime}

	return
}
