package storage

import (
	"errors"
	"time"

	"gopkg.in/gorp.v1"
)

// widget type
const (
	WidgetSpline = 1
	WidgetPie    = 2
	WidgetDonut  = 3
	WidgetNumber = 4
	WidgetColumn = 5
)

// A Widget is a view object to show data chart
type Widget struct {
	ID          int       `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DashboardID int       `db:"dashboard_id"`
	Config      string    `db:"config"`
	Type        int       `db:"type"`
}

func initWidgetTable(dbmap *gorp.DbMap) {
	widgetTable := dbmap.AddTableWithName(Widget{}, "widgets")
	widgetTable.SetKeys(true, "id")
}

// GetWidget will retrieve widget by widget id
func GetWidget(widgetID int, dbmap *gorp.DbMap) (widget *Widget) {
	err := dbmap.SelectOne(&widget, "select * from widgets where id = ?", widgetID)
	if err != nil {
		return nil
	}

	return
}

// QueryWidget will retrieve widgets by dashboard id
func QueryWidget(dashboardID int, dbmap *gorp.DbMap) (widgets []Widget, err error) {
	_, err = dbmap.Select(&widgets, "select * from widgets where dashboard_id = ?", dashboardID)
	return
}

// Save will create or insert database record
func (w *Widget) Save(dbmap *gorp.DbMap) (err error) {
	w.UpdatedAt = time.Now()

	var affectCount int64
	if w.ID > 0 {
		affectCount, err = dbmap.Update(w)
	} else {
		w.CreatedAt = w.UpdatedAt
		err = dbmap.Insert(w)
		affectCount = 1
	}

	if err != nil {
		return err
	}

	if affectCount == 0 {
		return errors.New("failed to save widget, affectCount = 0")
	}

	return
}

// Remove will remove database record
func (w *Widget) Remove(dbmap *gorp.DbMap) (err error) {
	_, err = dbmap.Delete(w)
	return
}
