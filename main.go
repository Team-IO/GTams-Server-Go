package main

import (
	"net/http"
	"./core"
	"./datastore"
)

func main() {

	core.InitLogging()

	core.Logger.Open()
	defer core.Logger.Close()
	core.Logger.Info("Staring up HTTP Server...")

	datastore.InitDatastore()
	defer datastore.CloseDatastore()

	core.Logger.Info("Staring up HTTP Server...")
	core.GenEndpoints(1)
	core.Logger.Info("HTTP Server started...")

	http.ListenAndServe(":60405", nil)

}