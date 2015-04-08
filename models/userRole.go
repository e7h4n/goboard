package models

import "gopkg.in/gorp.v1"

type UserRole struct {
	ID        int `db:"id"`
	UserID    int `db:"user_id"`
	RoleID    int `db:"role_id"`
	ProjectID int `db:"project_id"`
}

func initUserRoleTable(dbmap *gorp.DbMap) {
	userRoleTable := dbmap.AddTableWithName(UserRole{}, "user_roles")
	userRoleTable.SetKeys(true, "id")
}

func (ur *UserRole) Save(dbmap *gorp.DbMap) (err error) {
	return dbmap.Insert(ur)
}
