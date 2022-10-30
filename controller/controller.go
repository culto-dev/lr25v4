package controller

import (
	"lr25v4_back/utils"
	"net/http"
)

type LocationHandlerPair struct {
	Location string
	Handler  func(w http.ResponseWriter, r *http.Request)
}

func GetGoods(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()

	w.Write(utils.ObjectToJson(parameters))
}

func AllocateHandlers(pairs []LocationHandlerPair) {
	for _, pair := range pairs {
		http.HandleFunc(pair.Location, pair.Handler)
	}
}
