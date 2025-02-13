package main

import (
	"fmt"
	"greet/greet"
	"net/http"
)

const EN = "en"
const JA = "ja"
const KR = "kr"

const LOC = JA

func main() {
	var greeter greet.GreetInterface
	if LOC == EN {
		greeter = greet.NewGreetInEnglish()
	} else if LOC == JA {
		greeter = greet.NewGreetInJapanese()
	} else if LOC == KR {
		greeter = greet.NewGreetInHangeul()
	}

	// FOCUS: DI されている関数
	handler := Handler{greeter: greeter}

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

type Handler struct {
	greeter greet.GreetInterface
}

func (h *Handler) Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, h.greeter.Greet())
}

func (h *Handler) GreetWithName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintln(w, fmt.Sprintf("%s, %s", h.greeter.Greet(), name))
}
