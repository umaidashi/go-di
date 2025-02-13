package main

import (
	"fmt"
	"greet-no-di/greet"
	"net/http"
)

func main() {
	handler := Handler{}

	server := http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}

	fmt.Println("Server Started")
	server.ListenAndServe()
}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// greet.GreetInHangeul()
	// greet.GreetInJapanese()

	greet := greet.GreetInEnglish()

	fmt.Fprintln(w, greet)
}
