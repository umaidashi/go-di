package greet

type GreetInJapanese struct{}

func NewGreetInJapanese() GreetInterface {
	return &GreetInJapanese{}
}

func (g *GreetInJapanese) Greet() string {
	return "こんにちは!"
}
