package controller

import (
	"net/http"
)

type LocationHandlerPair struct {
	Location string
	Handler  func(w http.ResponseWriter, r *http.Request)
}

func AllocateHandlers(pairs []LocationHandlerPair) {
	for _, pair := range pairs {
		http.HandleFunc(pair.Location, pair.Handler)
	}
}
