package vo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yuantiku/goboard/storage"
)

func TestWidgetModel(t *testing.T) {
	dims := make([]WidgetDimensionConfig, 2)
	dims[0] = WidgetDimensionConfig{Key: "keyA", Name: "NameA", Value: "ValueA"}
	dims[1] = WidgetDimensionConfig{Key: "keyB", Name: "NameB", Value: "ValueB"}

	dataInfos := make([]WidgetDataSourceConfig, 2)
	dataInfos[0] = WidgetDataSourceConfig{ID: 1, Dimensions: dims}
	dataInfos[1] = WidgetDataSourceConfig{ID: 2}

	wc := WidgetConfig{Name: "Foo", Limit: 1, DataInfos: dataInfos}
	w := &Widget{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), Type: storage.WidgetPie, Config: wc}
	widget, err := w.Model()
	checkTestErr(err)

	assert.NotNil(t, widget)
	assert.Equal(t, w.ID, widget.ID)
	assert.Equal(t, w.Type, widget.Type)
	assert.Equal(t, w.CreatedAt, widget.CreatedAt)
	assert.Equal(t, w.UpdatedAt, widget.UpdatedAt)
	assert.Equal(t, w.DashboardID, widget.DashboardID)
	assert.Equal(t, `{"name":"Foo","limit":1,"dataInfos":[{"id":1,"dimensions":[{"key":"keyA","name":"NameA","value":"ValueA"},{"key":"keyB","name":"NameB","value":"ValueB"}]},{"id":2,"dimensions":null}]}`, widget.Config)
}

func TestNewWidget(t *testing.T) {
	widget := &storage.Widget{
		ID:          1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DashboardID: 1,
		Type:        storage.WidgetSpline,
		Config:      `{"name":"Foo","limit":1,"dataInfos":[{"id":1,"dimensions":[{"key":"keyA","name":"NameA","value":"ValueA"},{"key":"keyB","name":"NameB","value":"ValueB"}]},{"id":2,"dimensions":null}]}`}

	vo, err := NewWidget(widget)
	checkTestErr(err)
	assert.NotNil(t, vo)
	assert.Equal(t, "Foo", vo.Config.Name)
	assert.Len(t, vo.Config.DataInfos, 2)
}
