package greet

type GreetInEnglish struct{}

func NewGreetInEnglish() GreetInterface {
	return &GreetInEnglish{}
}

func (g *GreetInEnglish) Greet() string {
	return "Hello!"
}
