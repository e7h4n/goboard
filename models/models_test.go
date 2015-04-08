package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

func initTestDb(droptable bool) *gorp.DbMap {
	db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
	checkTestErr(err)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	InitDbMap(dbmap)

	if droptable {
		err = dbmap.DropTablesIfExists()
		checkTestErr(err)
	}

	err = dbmap.CreateTablesIfNotExists()
	checkTestErr(err)

	role := &Role{Name: "Admin", Scope: RoleGlobal}
	checkTestErr(role.Save(dbmap))

	return dbmap
}

func checkTestErr(err error) {
	if err != nil {
		panic(err)
	}
}
