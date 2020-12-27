package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	Request struct {
		Address  string `json:"address"`
		Password string `json:"password"`
		Command  string `json:"command"`
	}
)

func message(w http.ResponseWriter, req *http.Request) {

	b, err := ioutil.ReadAll(req.Body)

	defer req.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var r Request

	err = json.Unmarshal(b, &r)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, r.Command)
}

func main() {

	fs := http.FileServer(http.Dir("frontend/dist"))

	http.HandleFunc("/message", message)
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
