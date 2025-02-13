package greet

import "fmt"

type GreetInHangeul struct{}

func NewGreetInHangeul() GreetInterface {
	return &GreetInHangeul{}
}

func (g *GreetInHangeul) Greet() {
	fmt.Println("안녕하세요 세계")
}
