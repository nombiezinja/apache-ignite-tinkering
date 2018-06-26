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

	var DB *sql.DB
	var err error

	dbstring := fmt.Sprintf("Driver=Apache Ignite;ADDRESS=%s;Cache=%s",
		os.Getenv("address"), os.Getenv("cache"))

	DB, err = sql.Open("odbc", dbstring)
	u.FailOnError(err, "Failed to open ODBC connection")

	err = DB.Ping()
	u.FailOnError(err, "Failed to ping ODBC")

	fmt.Println("Apache Ignite connection successfully established")
}
