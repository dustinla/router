package handler

import (
	"encoding/json"
	"net/http"
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

func GetAllCustomers(writer http.ResponseWriter, request *http.Request) {
	if err := json.NewEncoder(writer).Encode(customers); err != nil {
		panic(err)
	}
}

func GetAllDeliver(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(deliver)

}
