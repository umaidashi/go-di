package main

import (
	"fmt"
	"greet-no-di/greet"
	"net/http"
)

const EN = "en"
const JA = "ja"
const KR = "kr"

const LOC = EN

func main() {
	handler := Handler{}

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	// curl http://localhost:8080/greet
	http.HandleFunc("/greet", handler.Greet)

	// curl http://localhost:8080/greetWithName\?name=hoge
	http.HandleFunc("/greetWithName", handler.GreetWithName)

	fmt.Println("Server Started")
	server.ListenAndServe()
}

type Handler struct{}

func (h *Handler) Greet(w http.ResponseWriter, r *http.Request) {
	var msg string
	if LOC == EN {
		msg = greet.GreetInEnglish()
	} else if LOC == JA {
		msg = greet.GreetInJapanese()
	} else if LOC == KR {
		msg = greet.GreetInHangeul()
	}

	fmt.Fprintln(w, msg)
}

func (h *Handler) GreetWithName(w http.ResponseWriter, r *http.Request) {
	var msg string
	if LOC == EN {
		msg = greet.GreetInEnglish()
	} else if LOC == JA {
		msg = greet.GreetInJapanese()
	} else if LOC == KR {
		msg = greet.GreetInHangeul()
	}

	name := r.URL.Query().Get("name")

	fmt.Fprintln(w, fmt.Sprintf("%s, %s", msg, name))
}
