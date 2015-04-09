package models

import (
	"gopkg.in/gorp.v1"
)

// RoleScope is user role scope type.
type RoleScope int

const (
	// RoleProject means this role only valid in specified project.
	RoleProject = 1
	// RoleGlobal means this role is valid in whole server.
	RoleGlobal = 2
)

// Role is user role type.
type Role struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Scope int    `db:"scope"`
}

func initRoleTable(dbmap *gorp.DbMap) {
	userTable := dbmap.AddTableWithName(Role{}, "roles")
	userTable.SetKeys(true, "id")
}

// GetAllRole will fetch all roles.
func GetAllRole(dbmap *gorp.DbMap) (roles []Role, err error) {
	_, err = dbmap.Select(&roles, "select * from roles")
	return
}

// QueryRoleByScope will fetch roles which match specified scope type.
func QueryRoleByScope(scope RoleScope, dbmap *gorp.DbMap) (roles []Role, err error) {
	_, err = dbmap.Select(&roles, "select * from roles where scope = ?", scope)
	return
}

// Save will insert a role record to database.
func (r *Role) Save(dbmap *gorp.DbMap) (err error) {
	return dbmap.Insert(r)
}
