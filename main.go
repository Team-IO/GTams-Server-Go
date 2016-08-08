package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"net/http"
	"io"
	"github.com/satori/go.uuid"
	"encoding/json"
	"./core"
)

var db, err = sql.Open("mysql", "gtams:gtams@/GTams")

type EAuthenticate struct {
	Token string	`json:"token"`
}

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

	core.Logger.Info("HTTP Server started...")

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