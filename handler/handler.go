package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Customer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}

type Deliver struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}

var customers = []Customer{
	{
		Name:    "A GmbH",
		Address: "A Strasse",
		Tel:     "A123123",
	},
	{
		Name:    "B GmbH",
		Address: "B Strasse",
		Tel:     "B123123123",
	},
}

var deliver = []Deliver{
	{
		Name:    "Lieferan",
		Address: "A Strasse",
		Tel:     "A123123",
	},
	{
		Name:    "Lieferant B",
		Address: "B Strasse",
		Tel:     "B123123123",
	},
}

var i = 0

func GetAllCustomers(writer http.ResponseWriter, request *http.Request) {

	setHeader(writer)
	if err := json.NewEncoder(writer).Encode(customers); err != nil {
		panic(err)
	}
}

func GetAllDeliver(writer http.ResponseWriter, request *http.Request) {
	setHeader(writer)
	json.NewEncoder(writer).Encode(deliver)
	i++
	println("Anzahl an Requests : " + strconv.Itoa(i))

}
func GetDeliverShow(writer http.ResponseWriter, request *http.Request) {
	setHeader(writer)
	vars := mux.Vars(request)
	id := vars["devId"]
	println(id)
	json.NewEncoder(writer).Encode(deliver)

}

func setHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
}
