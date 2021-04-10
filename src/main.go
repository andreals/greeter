package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Resposta struct {
	Status   string `json:"status"`
	Greeting string `json:"greeting"`
}

func main() {

	http.HandleFunc("/", HandlerGreeter)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func HandlerGreeter(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("hello")
	if name == "" {
		name = "Anonymous"
	}

	resp := Resposta{
		Status:   "OK",
		Greeting: fmt.Sprintf("Hello, %s", name),
	}

	if name == "there" {
		resp.Greeting = "General Kenobi"
	}

	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao codificar resposta :("))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}
