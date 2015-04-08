package models

import (
	"gopkg.in/gorp.v1"
)

func InitDbMap(dbmap *gorp.DbMap) {
	initUserTable(dbmap)
	initRoleTable(dbmap)
	initUserRoleTable(dbmap)
}
