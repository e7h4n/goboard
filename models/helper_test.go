package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

func testError(err error) {
	if err != nil {
		panic(err)
	}
}

func initDb(droptable bool) *gorp.DbMap {
	db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
	testError(err)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	initDbMap(dbmap)

	if droptable {
		err = dbmap.DropTablesIfExists()
		testError(err)
	}

	err = dbmap.CreateTablesIfNotExists()
	testError(err)

	return dbmap
}
