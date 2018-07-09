package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/alexbrainman/odbc"
	u "github.com/nombiezinja/chstub/utils"
)

func main() {
	Setenv()

	var db *sql.DB
	var err error

	dbstring := fmt.Sprintf("Driver=Apache Ignite;ADDRESS=%s;Cache=%s",
		os.Getenv("address"), os.Getenv("cache"))

	db, err = sql.Open("odbc", dbstring)
	u.FailOnError(err, "Failed to open ODBC connection")

	err = db.Ping()
	u.FailOnError(err, "Failed to ping ODBC")

	fmt.Println("Apache Ignite connection successfully established")

	created, err := db.Exec("CREATE TABLE FOOS (ID UUID PRIMARY KEY, CREATED_AT INT );")
	u.FailOnError(err, "Failed to create table")

	fmt.Println(created)

	inserted, err := db.Exec("INSERT INTO BUYS VALUES('a5595890-7f1e-403c-869b-70509a976e23', 'a5595890-7f1e-403c-869b-70509a976e23', 2.12, 2.12, 3.11, 1530124408)")
	u.FailOnError(err, "Failed to insert")

	fmt.Println(inserted)

	// db.Exec("INSERT INTO BUYS.BUY VALUES (1, 29, 29, 236, '2d931510-d99f-494a-8c67-87feb05e1594', 15003929329);")
}
