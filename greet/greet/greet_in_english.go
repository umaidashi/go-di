package greet

import "fmt"

type GreetInEnglish struct{}

func NewGreetInEnglish() GreetInterface {
	return &GreetInEnglish{}
}

func (g *GreetInEnglish) Greet() {
	fmt.Println("Hello World")
}
