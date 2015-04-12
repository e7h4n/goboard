package logical

import (
	"github.com/yuantiku/goboard/storage"
	"github.com/yuantiku/goboard/web/vo"
)

// GetWidget retrieve widget by id
func GetWidget(widgetID int, ctx *Context) (widget *vo.Widget, err error) {
	mWidget, err := storage.GetWidget(widgetID, ctx.DbMap)

	if err != nil {
		return nil, err
	}

	return vo.NewWidget(mWidget)
}

// QueryWidget retrieve widgets by dashboardID
func QueryWidget(dashboardID int, ctx *Context) (widgets []vo.Widget, err error) {
	mWidgets, err := storage.QueryWidget(dashboardID, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	widgets = make([]vo.Widget, len(mWidgets))
	for i, v := range mWidgets {
		widget, err := vo.NewWidget(&v)
		if err != nil {
			return nil, err
		}

		widgets[i] = *widget
	}

	return
}

// SaveWidget create or update widget to database
func SaveWidget(dashboardID int, widget *vo.Widget, ctx *Context) (err error) {
	widget.DashboardID = dashboardID
	mWidget, err := widget.Model()
	if err != nil {
		return err
	}

	err = mWidget.Save(ctx.DbMap)
	if err != nil {
		return err
	}

	savedWidget, err := vo.NewWidget(mWidget)
	if err != nil {
		return err
	}

	*widget = *savedWidget

	return
}

// RemoveWidget delete a widget
func RemoveWidget(widgetID int, ctx *Context) (err error) {
	widget := &storage.Widget{ID: widgetID}
	err = widget.Remove(ctx.DbMap)
	return
}
