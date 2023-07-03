package composit

type DuckFactory interface {
	createMallardDuck() Quackable
	createRedheadDuck() Quackable
	createDuckCall() Quackable
	createRubberDuck() Quackable
	createGooseAdapter() Quackable
}

type AbstractDuckFactory struct {
	DuckFactory
}

func (p *AbstractDuckFactory) createMallardDuck() Quackable {
	return NewQuackCounter(NewMallardDuck())
}
func (p *AbstractDuckFactory) createRedheadDuck() Quackable {
	return NewQuackCounter(NewRedheadDuck())
}
func (p *AbstractDuckFactory) createDuckCall() Quackable {
	return NewQuackCounter(NewDuckCall())
}
func (p *AbstractDuckFactory) createRubberDuck() Quackable {
	return NewQuackCounter(NewRubberDuck())
}
func (p *AbstractDuckFactory) createGooseAdapter() Quackable {
	return NewGooseAdapter(Goose{})
}
