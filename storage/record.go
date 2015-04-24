package storage

import (
	"time"

	"github.com/lann/squirrel"
	"gopkg.in/gorp.v1"
)

// Record stores data value and dim info
type Record struct {
	ID           int64     `db:"id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	DataSourceID int       `db:"data_source_id"`
	Value        float64   `db:"value"`
	Year         int       `db:"year"`
	Month        int       `db:"month"`
	Day          int       `db:"day"`
	Hour         int       `db:"hour"`
	Minute       int       `db:"minute"`
	Second       int       `db:"second"`
	DateTime     time.Time `db:"date_time"`
	Dim1         string    `db:"dim1"`
	Dim2         string    `db:"dim2"`
	Dim3         string    `db:"dim3"`
}

// RecordFilter is a paramete set to filter records for QueryRecord
type RecordFilter struct {
	Count     uint64
	Offset    uint64
	TimeLimit uint64
	StartTime time.Time
	EndTime   time.Time
	OrderBy   string
	Dim1      *string
	Dim2      *string
	Dim3      *string
}

// RecordDimFilter to make a dim filter variable for RecordFilter
func RecordDimFilter(dim string) *string {
	return &dim
}

func (rf RecordFilter) filterTimeSql(sqlBuilder squirrel.SelectBuilder) squirrel.SelectBuilder {
	if rf.StartTime.Unix() > 0 && rf.EndTime.Unix() > 0 {
		sqlBuilder = sqlBuilder.
			Where("date_time > ?", rf.StartTime).
			Where("date_time < ?", rf.EndTime)
	}

	if rf.Dim1 != nil {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"dim1": *rf.Dim1})
	}

	if rf.Dim2 != nil {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"dim2": *rf.Dim2})
	}

	if rf.Dim3 != nil {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"dim3": *rf.Dim3})
	}

	if rf.OrderBy != "" {
		sqlBuilder = sqlBuilder.OrderBy(rf.OrderBy + " desc")
	} else {
		sqlBuilder = sqlBuilder.OrderBy("date_time desc")
	}

	return sqlBuilder
}

func initRecordTable(dbmap *gorp.DbMap) {
	recordTable := dbmap.AddTableWithName(Record{}, "records")
	recordTable.SetKeys(true, "id")
	recordTable.ColMap("date_time").SetNotNull(true)
	recordTable.ColMap("data_source_id").SetNotNull(true)
	recordTable.ColMap("value").SetNotNull(true)
	recordTable.ColMap("created_at").SetNotNull(true)
	recordTable.ColMap("updated_at").SetNotNull(true)
	recordTable.SetUniqueTogether("data_source_id", "date_time", "dim1", "dim2", "dim3")
}

// GetRecord will retrieve by record id
func GetRecord(recordID int, dbmap *gorp.DbMap) (record *Record, err error) {
	err = dbmap.SelectOne(&record, "select * from records where id = ?", recordID)
	return
}

// QueryRecord will retrieve records by RecordFilter
func QueryRecord(dataSourceID int, filter RecordFilter, dbmap *gorp.DbMap) (records []Record, err error) {
	sqlQuery := squirrel.Select("*").From("records")
	sqlQuery = sqlQuery.Where(squirrel.Eq{"data_source_id": dataSourceID})

	if filter.Count > 0 {
		sqlQuery = sqlQuery.Limit(filter.Count)
	}

	if filter.Offset > 0 {
		sqlQuery = sqlQuery.Offset(filter.Offset)
	}

	sqlQuery = filter.filterTimeSql(sqlQuery)

	if filter.TimeLimit > 0 {
		type TimeRow struct {
			DateTime time.Time
		}

		timeQuery := squirrel.Select("distinct(date_time) as DateTime").From("records").
			Where(squirrel.Eq{"data_source_id": dataSourceID}).
			OrderBy("date_time desc").
			Limit(filter.TimeLimit)

		timeQuery = filter.filterTimeSql(timeQuery)

		timeSql, timeParams, err := timeQuery.ToSql()
		if err != nil {
			return nil, err
		}

		var timeRows []TimeRow
		_, err = dbmap.Select(&timeRows, timeSql, timeParams...)
		if err != nil {
			return nil, err
		}

		dateTimes := make([]time.Time, len(timeRows))
		for i, v := range timeRows {
			dateTimes[i] = v.DateTime
		}

		sqlQuery = sqlQuery.Where(squirrel.Eq{"date_time": dateTimes})
	}

	sql, params, err := sqlQuery.ToSql()

	_, err = dbmap.Select(&records, sql, params...)

	return
}

// Save will update or create a record
func (r *Record) Save(dbmap *gorp.DbMap) (err error) {
	localLoc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	r.UpdatedAt = time.Now()
	r.DateTime = time.Date(r.Year, time.Month(r.Month), r.Day, r.Hour, r.Minute, r.Second, 0, localLoc)

	if r.ID == 0 {
		r.ID, err = dbmap.SelectInt("select id from records where data_source_id = ? and date_time = ? and dim1 = ? and dim2 = ? and dim3 = ?", r.DataSourceID, r.DateTime, r.Dim1, r.Dim1, r.Dim3)
		if err != nil {
			return err
		}
	}

	if r.ID == 0 {
		r.CreatedAt = r.UpdatedAt
		err = dbmap.Insert(r)
		return
	}

	_, err = dbmap.Update(r)

	return
}

// Remove a record
func (r *Record) Remove(dbmap *gorp.DbMap) (err error) {
	_, err = dbmap.Delete(r)
	return
}

// ClearRecord remove all records belongs to specified data source
func ClearRecord(dataSourceID int, dbmap *gorp.DbMap) (err error) {
	_, err = dbmap.Exec("delete * from records where dataSourceID = ?", dataSourceID)
	return
}
