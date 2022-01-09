package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "henlo\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

type PassReq struct {
	Domain         string
	Username       string
	MasterPassword string
	Version        string
}

func pass(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var passReq PassReq
	err := decoder.Decode(&passReq)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// TODO: build password based on PassReq
	fmt.Fprintf(w, "Your password: %s", passReq.MasterPassword)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/pass", pass)
	http.ListenAndServe(":8090", nil)
}
