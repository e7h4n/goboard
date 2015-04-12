package logical

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuantiku/goboard/storage"
)

var userAdmin *storage.User
var userNorm *storage.User
var projectFoo *storage.Project
var projectBar *storage.Project

func TestAuthProject(t *testing.T) {
	ctx := initTest(true)

	userAdmin = &storage.User{Email: "admin@fenbi.com"}
	checkTestErr(userAdmin.Save(ctx.DbMap))

	userNorm = &storage.User{Email: "zhangyc@fenbi.com"}
	checkTestErr(userNorm.Save(ctx.DbMap))

	projectFoo = &storage.Project{Name: "project foo"}
	checkTestErr(projectFoo.Save(ctx.DbMap))

	projectBar = &storage.Project{Name: "project bar"}
	checkTestErr(projectBar.Save(ctx.DbMap))

	roles, err := storage.QueryRoleByScope(storage.RoleProject, ctx.DbMap)
	checkTestErr(err)
	roleID := roles[0].ID

	userRole := &storage.UserRole{UserID: userNorm.ID, ProjectID: projectFoo.ID, RoleID: roleID}
	checkTestErr(userRole.Save(ctx.DbMap))

	authorized, err := AuthProject(userAdmin.ID, projectFoo.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthProject(userAdmin.ID, projectBar.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthProject(userNorm.ID, projectFoo.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthProject(userNorm.ID, projectBar.ID, ctx)
	checkTestErr(err)
	assert.False(t, authorized)
}

func TestAuthDataSource(t *testing.T) {
	ctx := initTest(false)

	dataSourceFoo := &storage.DataSource{Name: "foo", ProjectID: projectFoo.ID, Key: "foo"}
	checkTestErr(dataSourceFoo.Save(ctx.DbMap))

	dataSourceBar := &storage.DataSource{Name: "bar", ProjectID: projectBar.ID, Key: "bar"}
	checkTestErr(dataSourceBar.Save(ctx.DbMap))

	authorized, err := AuthDataSource(userAdmin.ID, dataSourceFoo.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthDataSource(userAdmin.ID, dataSourceBar.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthDataSource(userNorm.ID, dataSourceFoo.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthDataSource(userNorm.ID, dataSourceBar.ID, ctx)
	checkTestErr(err)
	assert.False(t, authorized)
}

func TestAuthFolder(t *testing.T) {
	ctx := initTest(false)

	folderFoo := &storage.Folder{Name: "foo", ProjectID: projectFoo.ID}
	checkTestErr(folderFoo.Save(ctx.DbMap))

	folderBar := &storage.Folder{Name: "bar", ProjectID: projectBar.ID}
	checkTestErr(folderBar.Save(ctx.DbMap))

	authorized, err := AuthFolder(userAdmin.ID, folderFoo.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthFolder(userAdmin.ID, folderBar.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthFolder(userNorm.ID, folderFoo.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthFolder(userNorm.ID, folderBar.ID, ctx)
	checkTestErr(err)
	assert.False(t, authorized)
}

func TestAuthDashboard(t *testing.T) {
	ctx := initTest(false)

	dashboardFoo := &storage.Dashboard{Name: "foo", ProjectID: projectFoo.ID}
	checkTestErr(dashboardFoo.Save(ctx.DbMap))

	dashboardBar := &storage.Dashboard{Name: "bar", ProjectID: projectBar.ID}
	checkTestErr(dashboardBar.Save(ctx.DbMap))

	authorized, err := AuthDashboard(userAdmin.ID, dashboardFoo.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthDashboard(userAdmin.ID, dashboardBar.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthDashboard(userNorm.ID, dashboardFoo.ID, ctx)
	checkTestErr(err)
	assert.True(t, authorized)

	authorized, err = AuthDashboard(userNorm.ID, dashboardBar.ID, ctx)
	checkTestErr(err)
	assert.False(t, authorized)
}
