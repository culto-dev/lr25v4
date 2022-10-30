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
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	w.Write(utils.ObjectToJson([]byte(limit + offset)))
}

func GetQueryParams(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	w.Write(utils.ObjectToJson(params))
}

func AllocateHandlers(pairs []LocationHandlerPair) {
	for _, pair := range pairs {
		http.HandleFunc(pair.Location, pair.Handler)
	}
}
