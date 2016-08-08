package sys

import (
	"net/http"
	"github.com/satori/go.uuid"
	"encoding/json"
	"./../../core"
)

type AuthenticateResponse struct {
	Token string	`json:"token"`
}

func Auth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := AuthenticateResponse{
		Token: uuid.NewV4().String(),
	}
	//db.Exec("INSERT INTO installation (uuid) VALUES ('" + response.Token + "');")
	core.Logger.Info("Authenticating Client with 'token' random UUID: " + response.Token)
	w.WriteHeader(http.StatusCreated)
	data, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.Write(data)
}