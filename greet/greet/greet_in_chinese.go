package greet

type GreetInChinese struct{}

func NewGreetInChinese() GreetInterface {
	return &GreetInChinese{}
}

func (g *GreetInChinese) Greet() string {
	return "你好!"
}
