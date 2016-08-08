package endpoints

import (
	"../handlers/sys"
	"../handlers/player"
	"../handlers/market"
	"../handlers/terminal"
	"net/http"
	"fmt"
)

type Endpoint struct {
	Endpoint string
	Handler  func(http.ResponseWriter, *http.Request)
	Version  int
	Id       int
}

func GenEndpoints(version int) {
	Endpoints := []Endpoint{
		{
			Endpoint: "/authenticate",
			Handler: sys.Auth,
			Version: 1,
			Id: 1,
		},
		{
			Endpoint: "/terminal/new",
			Handler: terminal.Add,
			Version: 1,
			Id: 2,
		},
		{
			Endpoint: "/terminal/destroy",
			Handler: terminal.Destroy,
			Version: 1,
			Id: 3,
		},
		{
			Endpoint: "/terminal/status",
			Handler: terminal.Status,
			Version: 1,
			Id: 4,
		},
		{
			Endpoint: "/terminal/owner",
			Handler: terminal.Owner,
			Version: 1,
			Id: 5,
		},
		{
			Endpoint: "/terminal/trades",
			Handler: terminal.Trades,
			Version: 1,
			Id: 6,
		},
		{
			Endpoint: "/terminal/trades/add",
			Handler: terminal.AddTrade,
			Version: 1,
			Id: 7,
		},
		{
			Endpoint: "/terminal/trades/remove",
			Handler: terminal.RemoveTrade,
			Version: 1,
			Id: 8,
		},
		{
			Endpoint: "/terminal/goods",
			Handler: terminal.GoodsGet,
			Version: 1,
			Id: 9,
		},
		{
			Endpoint:"/terminal/goods/add",
			Handler: terminal.GoodsAdd,
			Version: 1,
			Id: 10,
		},
		{
			Endpoint: "/terminal/goods/remove",
			Handler: terminal.GoodsRemove,
			Version: 1,
			Id: 11,
		},
		{
			Endpoint: "/player/status",
			Handler: player.Status,
			Version: 1,
			Id: 12,
		},
		{
			Endpoint: "/market/query",
			Handler: market.Info,
			Version: 1,
			Id: 13,
		},

	}
	for i := 0; i < len(Endpoints); i++ {
		if Endpoints[i].Version == version {
			http.HandleFunc("/V" + fmt.Sprintf("%d", Endpoints[i].Version) + Endpoints[i].Endpoint, Endpoints[i].Handler)
		}
	}
}