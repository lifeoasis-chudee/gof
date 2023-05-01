package starbuzz

import "fmt"

type Espresso struct {
	size Size
}

func NewEspresso() Beverage {
	return &Espresso{}
}

func (e *Espresso) getDescription() string {
	return fmt.Sprintf("에스프레소 커피 %s", e.size)
}

func (e *Espresso) cost() float32 {
	return 1.99
}

func (e *Espresso) setSize(size Size) {
	e.size = size
}

func (e *Espresso) getSize() Size {
	return e.size
}
