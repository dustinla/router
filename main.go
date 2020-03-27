package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
)

type url struct {
	Url string `json:"url"`
}

func main() {
	//StartServer("1234")

	test := url{Url: "http://polyglot.ninja/golang-making-http-requests/"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(test)
	resp, err := http.Post("https://rel.ink/api/links/", "application/json", b)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	str3 := bytes.NewBuffer(bodyBytes).String()
	res := gjson.Get(str3, "url")

	fmt.Println(res)
}
