package main

import (
	"greet/greet"
)

func main() {
	greet := greet.NewGreetInHangeul()
	// greet := greet.NewGreetInJapanese()
	// greet := greet.NewGreetInEnglish()

	// FOCUS: DI されている関数
	greeter(greet)
}

func greeter(g greet.GreetInterface) {
	g.Greet()
}
