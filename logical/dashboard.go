package logical

import (
	"github.com/yuantiku/goboard/storage"
	"github.com/yuantiku/goboard/web/vo"
)

// QueryDashboard will retrieve dashboards by project and user, show all public and user's private board
func QueryDashboard(projectID int, ctx *Context) (dashboards []vo.Dashboard, err error) {
	mDashboards, err := storage.QueryDashboard(projectID, ctx.UserID, ctx.DbMap)

	dashboards = make([]vo.Dashboard, len(mDashboards))
	for i, v := range mDashboards {
		dashboard, err := vo.NewDashboard(&v)
		if err != nil {
			return nil, err
		}

		dashboards[i] = *dashboard
	}

	return
}

// GetDashboard retrieve dashboard by id
func GetDashboard(dashboardID int, ctx *Context) (dashboard *vo.Dashboard, err error) {
	mDashboard, err := storage.GetDashboard(dashboardID, ctx.UserID, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	dashboard, err = vo.NewDashboard(mDashboard)
	if err != nil {
		return nil, err
	}

	return
}

// SaveDashboard will save or update dashboard
func SaveDashboard(projectID int, dashboard *vo.Dashboard, ctx *Context) (err error) {
	if dashboard.ID > 0 {
		mDashboard, err := storage.GetDashboard(dashboard.ID, ctx.UserID, ctx.DbMap)
		if err != nil {
			return err
		}

		dashboard.OwnerID = mDashboard.OwnerID
		dashboard.ProjectID = mDashboard.ProjectID

		// only owner can change private attribute
		if dashboard.OwnerID != ctx.UserID {
			dashboard.Private = mDashboard.Private
		}
	} else {
		dashboard.OwnerID = ctx.UserID
		dashboard.ProjectID = projectID
	}

	mDashboard, err := dashboard.Model()
	if err != nil {
		return err
	}

	err = mDashboard.Save(ctx.DbMap)
	if err != nil {
		return err
	}

	newDashboard, err := vo.NewDashboard(mDashboard)
	if err != nil {
		return err
	}

	*dashboard = *newDashboard

	return
}

// RemoveDashboard delete a dashboard
func RemoveDashboard(dashboardID int, ctx *Context) (err error) {
	dashboard, err := storage.GetDashboard(dashboardID, ctx.UserID, ctx.DbMap)
	if err != nil {
		return err
	}

	err = dashboard.Remove(ctx.DbMap)
	return
}
