package logical

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuantiku/goboard/web/vo"
)

func TestSaveRecord(t *testing.T) {
	ctx := initTest(true)

	dataSource := &vo.DataSource{Name: "ds"}
	dataSource.Config.Dimensions = make([]vo.DimensionConfig, 2)
	dataSource.Config.Dimensions[0] = vo.DimensionConfig{Key: "hello"}
	dataSource.Config.Dimensions[1] = vo.DimensionConfig{Key: "world"}

	err := SaveDataSource(1, dataSource, ctx)
	checkTestErr(err)

	record := &vo.Record{DataSourceID: 1, Dim1: "foo", Dim2: "bar"}
	err = SaveRecord(record, ctx)
	checkTestErr(err)

	assert.Equal(t, 1, record.ID)
}

func TestGetRecord(t *testing.T) {
	ctx := initTest(false)

	record, err := GetRecord(1, ctx)
	checkTestErr(err)
	assert.Equal(t, "foo", record.Dim1)
}

func TestRecrodToJSON(t *testing.T) {
	ctx := initTest(false)

	record, err := GetRecord(1, ctx)
	checkTestErr(err)

	jRecord, err := RecordToJSON(record, ctx)
	checkTestErr(err)

	assert.Equal(t, `{"dataSourceId":1,"dateTime":"0001-01-01T00:00:00Z","day":0,"dim3":"","hello":"foo","hour":0,"id":1,"minute":0,"month":0,"second":0,"value":0,"world":"bar","year":0}`, jRecord)
}

func TestRecordFromJSON(t *testing.T) {
	ctx := initTest(false)

	jRecord := `{"dataSourceId":1,"dateTime":"0001-01-01T00:00:00Z","day":0,"dim3":"","hello":"foo","hour":0,"id":1,"minute":0,"month":0,"second":0,"value":0,"world":"bar","year":0}`
	record, err := RecrodFromJSON(jRecord, ctx)
	checkTestErr(err)

	assert.Equal(t, "foo", record.Dim1)
}
