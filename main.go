package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"plugin"
)

type Task struct {
	Type string
	Args []string
}

func handler(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var task Task
	err := decoder.Decode(&task)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	p, err := plugin.Open("plugins/" + task.Type + ".so")

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := p.Lookup("Process")

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := f.(func(args ...string) ([]byte, error))(task.Args...)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(resp)
}

func main() {
	fmt.Println("starting")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
