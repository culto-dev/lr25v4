package server

import (
	"fmt"
	"lr25v4_back/config"
	"lr25v4_back/controller"
	"lr25v4_back/utils"
	"net/http"
)

// also works as below, but returns json keys capitalised
//
// type Animal struct {
//	Name string
//	Says string
// }

type Animal struct {
	Name string `json:"name"`
	Says string `json:"says"`
}

func handleSimple(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json")
	fmt.Println(r)
	w.Write(utils.ObjectToJson(Animal {Name: "goat", Says: "mleigh"}))
}

func StartServer() {
	var pairs []controller.LocationHandlerPair
	// wrong queries fallback to /
	// pairs = append(pairs, controller.LocationHandlerPair{ Location: "/", Handler: handleSimple })
	pairs = append(pairs, controller.LocationHandlerPair{ Location: "/gog", Handler: controller.GetGoods })
	pairs = append(pairs, controller.LocationHandlerPair{ Location: "/return-query-params", Handler: controller.GetQueryParams })
	controller.AllocateHandlers(pairs)
	http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), nil)
}
