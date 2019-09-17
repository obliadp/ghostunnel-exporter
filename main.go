package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var Keys []Metric

type Metric struct {
	Value  interface{} `json:"value"`
	Metric string      `json:"metric"`
}

func hello(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		t := strings.NewReplacer(".", " ", "-", " ")
		r := strings.NewReplacer(".", "_", "-", "_")
		for _, v := range Keys {
			fmt.Fprintf(w, "# %s\n", t.Replace(strings.Title(v.Metric)))
			fmt.Fprintf(w, "%s: %v\n", r.Replace(v.Metric), v.Value)
		}

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(body, &Keys); err != nil {
			panic(err)
		}
		r.Body.Close()

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func main() {

	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
