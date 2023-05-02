package starbuzz

import "fmt"

type Espresso struct {
	size    Size
	options map[string]int32
}

func NewEspresso() Beverage {
	return &Espresso{
		size:    TALL,
		options: make(map[string]int32),
	}
}

func (b *Espresso) getDescription() string {
	return fmt.Sprintf("에스프레소 커피 %s", b.size)
}

func (b *Espresso) cost() float32 {
	return 1.99
}

func (b *Espresso) setSize(size Size) {
	b.size = size
}

func (b *Espresso) getSize() Size {
	return b.size
}

func (b *Espresso) getOptions() map[string]int32 {
	return b.options
}

func (b *Espresso) setOptions(name string) {
	b.options[name] = b.options[name] + 1
}
