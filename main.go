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

	// created, err := db.Exec("CREATE TABLE TESTTEST(ID INT PRIMARY KEY, NAME VARCHAR(255));")
	// u.FailOnError(err, "Failed to create table")

	// fmt.Println(created)

	// inserted, err := db.Exec("INSERT INTO TESTTEST VALUES(17, 'BLEH')")
	// u.FailOnError(err, "Failed to insert")

	// fmt.Println(inserted)

	// db.Exec("INSERT INTO Organization.Organization (name) VALUES ('Hp');")
	db.Exec("INSERT INTO Buys.BUY VALUES (1, 29, 29, 236, '2d931510-d99f-494a-8c67-87feb05e1594', 15003929329);")
}
