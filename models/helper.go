package models

import "gopkg.in/gorp.v1"

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func initDbMap(dbmap *gorp.DbMap) {
	initUserTable(dbmap)
}
