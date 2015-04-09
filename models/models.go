package models

import (
	"gopkg.in/gorp.v1"
)

// InitDbMap will initialize database table relationships
func InitDbMap(dbmap *gorp.DbMap) {
	initUserTable(dbmap)
	initRoleTable(dbmap)
	initUserRoleTable(dbmap)
	initProjectTable(dbmap)
	initRolePrivilegeTable(dbmap)
	initDashboardTable(dbmap)
	initDataSourceTable(dbmap)
	initFolderTable(dbmap)
	initWidgetTable(dbmap)
}

// InitPrivilegeData will initialize privilege data to database
func InitPrivilegeData(dbmap *gorp.DbMap) (err error) {
	// initialize roles
	role := &Role{Name: "Admin", Scope: RoleGlobal}
	if err = role.Save(dbmap); err != nil {
		return
	}
	adminRoleID := role.ID

	role = &Role{Name: "User", Scope: RoleProject}
	if err = role.Save(dbmap); err != nil {
		return
	}
	userRoleID := role.ID

	// initialize admin privilege
	rolePrivilege := &RolePrivilege{Resource: ResourceProject, Operation: OperationGet, RoleID: adminRoleID}
	if err = rolePrivilege.Save(dbmap); err != nil {
		return
	}

	rolePrivilege = &RolePrivilege{Resource: ResourceProject, Operation: OperationPut, RoleID: adminRoleID}
	if err = rolePrivilege.Save(dbmap); err != nil {
		return
	}

	rolePrivilege = &RolePrivilege{Resource: ResourceProject, Operation: OperationPost, RoleID: adminRoleID}
	if err = rolePrivilege.Save(dbmap); err != nil {
		return
	}

	rolePrivilege = &RolePrivilege{Resource: ResourceProject, Operation: OperationDelete, RoleID: adminRoleID}
	if err = rolePrivilege.Save(dbmap); err != nil {
		return
	}

	// initialize user privilege
	rolePrivilege = &RolePrivilege{Resource: ResourceProject, Operation: OperationGet, RoleID: userRoleID}
	if err = rolePrivilege.Save(dbmap); err != nil {
		return
	}

	rolePrivilege = &RolePrivilege{Resource: ResourceProject, Operation: OperationPut, RoleID: userRoleID}
	if err = rolePrivilege.Save(dbmap); err != nil {
		return
	}

	return
}
