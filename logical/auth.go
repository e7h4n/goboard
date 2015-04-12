package logical

import (
	"github.com/yuantiku/goboard/storage"
)

// AuthProject check user privilege to specified project
func AuthProject(userID int, projectID int, ctx *Context) (authorized bool, err error) {
	return storage.AuthProject(userID, storage.ResourceProject, storage.OperationGet, projectID, ctx.DbMap)
}

// AuthDataSource check user privilege to specified data source
func AuthDataSource(userID int, dataSourceID int, ctx *Context) (authorized bool, err error) {
	dataSource, err := storage.GetDataSource(dataSourceID, ctx.DbMap)
	if err != nil {
		return false, err
	}

	return AuthProject(userID, dataSource.ProjectID, ctx)
}

// AuthFolder check user privilege to specified folder
func AuthFolder(userID int, folderID int, ctx *Context) (authorized bool, err error) {
	folder, err := storage.GetFolder(folderID, ctx.DbMap)
	if err != nil {
		return false, err
	}

	return AuthProject(userID, folder.ProjectID, ctx)
}

// AuthDashboard check user privilege to specified dashboard
func AuthDashboard(userID int, dashboardID int, ctx *Context) (authorized bool, err error) {
	dashboard, err := storage.GetDashboard(dashboardID, ctx.DbMap)
	if err != nil {
		return false, err
	}

	if dashboard.OwnerID != userID && dashboard.Private {
		return false, nil
	}

	return AuthProject(userID, dashboard.ProjectID, ctx)
}
