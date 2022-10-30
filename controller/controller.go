package controller

import (
	"lr25v4_back/service"
	"lr25v4_back/utils"
	"net/http"
	"strconv"
)

type LocationHandlerPair struct {
	Location string
	Handler  func(w http.ResponseWriter, r *http.Request)
}

func GetGoods(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	w.Write(utils.ObjectToJson(service.GetGoods(limit, offset)))
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
