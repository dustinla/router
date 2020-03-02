package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/deliver", getAllDeliver)
	router.HandleFunc("/test", getTest)
	log.Println("CustomerServer: Listening on " +
		"http://localhost:8080/customers ...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getTest(writer http.ResponseWriter, request *http.Request) {

}

