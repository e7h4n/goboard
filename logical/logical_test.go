package logical

import (
	"database/sql"
	"log"
	"os"

	"github.com/yuantiku/goboard/storage"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

func initTest(droptable bool) *Context {
	db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
	checkTestErr(err)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	storage.InitDbMap(dbmap)

	if droptable {
		err = dbmap.DropTablesIfExists()
		checkTestErr(err)

		err = dbmap.CreateTables()
		checkTestErr(err)

		err = storage.InitPrivilegeData(dbmap)
		checkTestErr(err)
	}

	ctx := &Context{DbMap: dbmap}

	return ctx
}

func enableTrace(ctx *Context) {
	ctx.DbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
}

func checkTestErr(err error) {
	if err != nil {
		panic(err)
	}
}
