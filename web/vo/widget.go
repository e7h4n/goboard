package vo

import (
	"encoding/json"
	"time"

	"github.com/yuantiku/goboard/storage"
)

// Widget is view object for storage.Widget
type Widget struct {
	ID          int          `json:"id"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DashboardID int          `json:"dashboardId"`
	Type        int          `json:"type"`
	Config      WidgetConfig `json:"config"`
}

// WidgetConfig is config for widget
type WidgetConfig struct {
	Name      string                   `json:"name"`
	Limit     int                      `json:"limit"`
	DataInfos []WidgetDataSourceConfig `json:"dataInfos"`
}

// WidgetDataSourceConfig is data source config for widget
type WidgetDataSourceConfig struct {
	ID         int                     `json:"id"`
	Dimensions []WidgetDimensionConfig `json:"dimensions"`
}

// WidgetDimensionConfig is widget dimension config
type WidgetDimensionConfig struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Model convert vo to storage model
func (w *Widget) Model() (widget *storage.Widget, err error) {
	widget = &storage.Widget{ID: w.ID, CreatedAt: w.CreatedAt, UpdatedAt: w.UpdatedAt, DashboardID: w.DashboardID, Type: w.Type}
	config, err := json.Marshal(w.Config)
	if err != nil {
		return
	}

	widget.Config = string(config)
	return
}

// NewWidget convert storage model to vo
func NewWidget(w *storage.Widget) (widget *Widget, err error) {
	widget = &Widget{}
	widget.ID = w.ID
	widget.CreatedAt = w.CreatedAt
	widget.UpdatedAt = w.UpdatedAt
	widget.DashboardID = w.DashboardID
	widget.Type = w.Type
	err = json.Unmarshal([]byte(w.Config), &widget.Config)

	return
}
