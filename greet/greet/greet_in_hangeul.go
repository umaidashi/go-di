package greet

type GreetInHangeul struct{}

func NewGreetInHangeul() GreetInterface {
	return &GreetInHangeul{}
}

func (g *GreetInHangeul) Greet() string {
	return "안녕하세요 세계"
}
