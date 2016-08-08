package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"net/http"
	"io"
	"github.com/satori/go.uuid"
	"github.com/go-ozzo/ozzo-log"
	"encoding/json"
)

var logger = log.NewLogger()
var db, err = sql.Open("mysql", "gtams:gtams@/GTams")

type EAuthenticate struct {
	Token string	`json:"token"`
}

func main() {

	targetLogFile := log.NewFileTarget()
	targetLogConsole := log.NewConsoleTarget()
	targetLogFile.FileName = "GTams-Server.log"
	targetLogFile.BackupCount = 2
	//targetLogConsole.MaxLevel = log.LevelError
	targetLogConsole.ColorMode = true
	logger.Targets = append(logger.Targets, targetLogConsole, targetLogFile)

	logger.Open()
	defer logger.Close()

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	logger.Info("Staring up HTTP Server...")

	http.HandleFunc("/authenticate", Auth)
	/*http.HandleFunc("/terminal/new", NewTerminal)
	http.HandleFunc("/terminal/destroy", DestroyTermial)
	http.HandleFunc("/terminal/status", TerminalStatus)
	http.HandleFunc("/terminal/owner", TerminalOwner)

	http.HandleFunc("/terminal/trades",  TerminalTrades)
	http.HandleFunc("/terminal/trades/add",  TerminalTrade)
	http.HandleFunc("/terminal/trades/remove",  TerminalRemoveTrade)
	
	http.HandleFunc("/terminal/goods",  TerminalGoodsGet)
	http.HandleFunc("/terminal/goods/add",  TerminalGoodsAdd)
	http.HandleFunc("/terminal/goods/remove",  TerminalGoodsRemove)
	
	http.HandleFunc("/player/status",  PlayerStatus)
	http.HandleFunc("/market/query",  MarketInfo)*/

	logger.Info("HTTP Server started...")

	http.ListenAndServe(":60405", nil)

}

func Auth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := EAuthenticate{
		Token: uuid.NewV4().String(),
	}
	db.Exec("INSERT INTO installation (uuid) VALUES ('" + response.Token + "');")
	logger.Info("Authenticating Client with 'token' random UUID: " + response.Token)
	w.WriteHeader(http.StatusCreated)
	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(response.Token)
	w.Write(data)
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	io.WriteString(w, "Hello World\n")
}