package storage

import (
	"database/sql"
	_ "log"
	_ "os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

func initTestDb(droptable bool) *gorp.DbMap {
	db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
	checkTestErr(err)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	InitDbMap(dbmap)

	if droptable {
		err = dbmap.DropTablesIfExists()
		checkTestErr(err)

		err = dbmap.CreateTables()
		checkTestErr(err)

		err = InitPrivilegeData(dbmap)
		checkTestErr(err)
	}

	return dbmap
}

func checkTestErr(err error) {
	if err != nil {
		panic(err)
	}
}
