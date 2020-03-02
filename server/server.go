package server

import (
	"github.com/dustinla/router/handler"
	"log"
	"net/http"
)

func StartServer(port string) {
	router := http.NewServeMux()
	router.HandleFunc("/customers", handler.GetAllCustomers)
	router.HandleFunc("/deliver", handler.GetAllDeliver)
	log.Printf("CustomerServer: Listening on "+
		"http://localhost:%v/customers ...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
