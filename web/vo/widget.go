package vo

import (
	"encoding/json"
	"time"

	"github.com/yuantiku/goboard/storage"
)

type Widget struct {
	ID          int          `json:"id"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DashboardID int          `json:"dashboardId"`
	Type        int          `json:"type"`
	Config      WidgetConfig `json:"config"`
}

type WidgetConfig struct {
	Name      string                   `json:"name"`
	Limit     int                      `json:"limit"`
	DataInfos []WidgetDataSourceConfig `json:"dataInfos"`
}

type WidgetDataSourceConfig struct {
	ID         int                     `json:"id"`
	Dimensions []WidgetDimensionConfig `json:"dimensions"`
}

type WidgetDimensionConfig struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (w *Widget) Model() (widget *storage.Widget, err error) {
	widget = &storage.Widget{ID: w.ID, CreatedAt: w.CreatedAt, UpdatedAt: w.UpdatedAt, DashboardID: w.DashboardID, Type: w.Type}
	config, err := json.Marshal(w.Config)
	if err != nil {
		return
	}

	widget.Config = string(config)
	return
}
