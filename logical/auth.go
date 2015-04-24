package logical

import (
	"github.com/perfectworks/goboard/storage"
)

// AuthProject check user privilege to specified project
func AuthProject(projectID int, ctx *Context) (authorized bool, err error) {
	return storage.AuthProject(ctx.UserID, storage.ResourceProject, storage.OperationGet, projectID, ctx.DbMap)
}

// AuthDataSource check user privilege to specified data source
func AuthDataSource(dataSourceID int, ctx *Context) (authorized bool, err error) {
	dataSource, err := storage.GetDataSource(dataSourceID, ctx.DbMap)
	if err != nil {
		return false, err
	}

	return AuthProject(dataSource.ProjectID, ctx)
}

// AuthFolder check user privilege to specified folder
func AuthFolder(folderID int, ctx *Context) (authorized bool, err error) {
	folder, err := storage.GetFolder(folderID, ctx.DbMap)
	if err != nil {
		return false, err
	}

	return AuthProject(folder.ProjectID, ctx)
}

// AuthDashboard check user privilege to specified dashboard
func AuthDashboard(dashboardID int, ctx *Context) (authorized bool, err error) {
	dashboard, err := storage.GetDashboard(dashboardID, ctx.UserID, ctx.DbMap)
	if err != nil {
		return false, err
	}

	if dashboard.OwnerID != ctx.UserID && dashboard.Private {
		return false, nil
	}

	return AuthProject(dashboard.ProjectID, ctx)
}
