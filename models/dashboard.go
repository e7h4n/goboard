package models

import (
	"errors"
	"time"

	"gopkg.in/gorp.v1"
)

// A Dashboard is a container for widgets
type Dashboard struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string    `db:"name"`
	Config    string    `db:"config"`
	ProjectID int       `db:"project_id"`
	Private   bool      `db:"private"`
	OwnerID   int       `db:"owner_id"`
}

func initDashboardTable(dbmap *gorp.DbMap) {
	dashboardTable := dbmap.AddTableWithName(Dashboard{}, "dashboards")
	dashboardTable.SetKeys(true, "id")
}

// GetDashboardByID will retrieve specified dashboard by id.
func GetDashboardByID(dashboardID int, dbmap *gorp.DbMap) (dashboard *Dashboard) {
	err := dbmap.SelectOne(&dashboard, "select * from dashboards where id = ?", dashboardID)
	if err != nil {
		return nil
	}

	return
}

/*
QueryDashboardByUser will retrieve such dashboards:

  1. all public boards which belongs to specified project
  2. all private boards which belongs to specified project and specified user
*/
func QueryDashboardByUser(projectID int, ownerID int, dbmap *gorp.DbMap) (dashboards []Dashboard, err error) {
	_, err = dbmap.Select(&dashboards, "select * from dashboards"+
		" where project_id = ?"+
		" and (owner_id = ? or private = ?)", projectID, ownerID, false)
	return
}

// Save will insert or update a dashboard record to database.
func (d *Dashboard) Save(dbmap *gorp.DbMap) (err error) {
	var affectCount int64
	d.UpdatedAt = time.Now()

	if d.ID > 0 {
		affectCount, err = dbmap.Update(d)
	} else {
		d.CreatedAt = d.UpdatedAt
		err = dbmap.Insert(d)
		affectCount = 1
	}

	if err != nil {
		return err
	}

	if affectCount == 0 {
		return errors.New("failed to save dashboard, affectCount = 0")
	}

	return
}

// Remove a dashboard record.
func (d *Dashboard) Remove(dbmap *gorp.DbMap) (err error) {
	_, err = dbmap.Delete(d)
	return
}
