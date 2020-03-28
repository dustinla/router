package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/dustinla/router/server"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type url struct {
	Url string `json:"url"`
}

func main() {
	go exit()
	go loop(1)
	go loop(2)
	go loop(3)

	StartServer("1234")

	test := url{Url: "http://polyglot.ninja/golang-making-http-requests/"}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(test)
	if err != nil {
	}
	resp, err := http.Post("https://rel.ink/api/links/", "application/json", b)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	str3 := bytes.NewBuffer(bodyBytes).String()
	res := gjson.Get(str3, "hashid")

	fmt.Println("https://rel.ink/" + res.String())

}

func loop(ch int) {

	for i := 0; ; i++ {

		blaa, _ := http.Get("http://localhost:1234/deliver")
		bli, err := ioutil.ReadAll(blaa.Body)
		if err != nil {
			os.Exit(3)
		}

		println(string(bli))
		chann := strconv.Itoa(ch)
		println(strconv.Itoa(i) + " : " + chann)
	}
}
func exit() {

	time.Sleep(5 * time.Second)
	os.Exit(0)
}
