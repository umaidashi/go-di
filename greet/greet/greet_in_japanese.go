package greet

import "fmt"

type GreetInJapanese struct{}

func NewGreetInJapanese() GreetInterface {
	return &GreetInJapanese{}
}

func (g *GreetInJapanese) Greet() {
	fmt.Println("こんにちは世界")
}
