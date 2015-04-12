package logic

import (
	"encoding/json"

	"github.com/yuantiku/goboard/storage"
	"github.com/yuantiku/goboard/web/vo"
)

func WrapWidget(w *storage.Widget) (widget *vo.Widget) {
	widget.ID = w.ID
	widget.CreatedAt = w.CreatedAt
	widget.UpdatedAt = w.UpdatedAt
	widget.DashboardID = w.DashboardID
	widget.Type = w.Type
	json.Unmarshal([]byte(w.Config), &widget.Config)

	return widget
}
