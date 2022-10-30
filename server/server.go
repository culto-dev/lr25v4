package server

import (
	"fmt"
	"lr25v4_back/controller"
	"lr25v4_back/utils"
	"net/http"
)

// also works as below, but returns json keys capitalised

//type Animal struct {
//	Name string
//	Says string
//}

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
	pairs = append(pairs, controller.LocationHandlerPair{ Location: "/", Handler: handleSimple })
	controller.AllocateHandlers(pairs)
	http.ListenAndServe(":8000", nil)
}
