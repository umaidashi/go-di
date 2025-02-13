package main

import (
	"fmt"
	"greet/greet"
	"net/http"
)

func main() {
	// greet := greet.NewGreetInHangeul()
	greet := greet.NewGreetInJapanese()
	// greet := greet.NewGreetInEnglish()

	// FOCUS: DI されている関数
	handler := Handler{greeter: greet}

	server := http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}

	fmt.Println("Server Started")
	server.ListenAndServe()
}

type Handler struct {
	greeter greet.GreetInterface
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, h.greeter.Greet())
}
