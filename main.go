package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"net/http"
	"./core"
)

var db, err = sql.Open("mysql", "gtams:gtams@/GTams")



func main() {

	core.InitLogging()

	core.Logger.Open()
	defer core.Logger.Close()

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	core.Logger.Info("Staring up HTTP Server...")
	core.GenEndpoints(1)
	core.Logger.Info("HTTP Server started...")

	http.ListenAndServe(":60405", nil)

}



