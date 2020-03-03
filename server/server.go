package server

import (
	"github.com/dustinla/router/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customers", handler.GetAllCustomers)
	router.HandleFunc("/deliver", handler.GetAllDeliver)
	router.HandleFunc("/deliver/{devId}", handler.GetDeliverShow)
	log.Printf("CustomerServer: Listening on "+
		"http://localhost:%v/customers ...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
