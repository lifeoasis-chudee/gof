package starbuzz

import "fmt"

type Decaf struct {
	size    Size
	options map[string]int32
}

func NewDecaf() Beverage {
	return &Decaf{
		size:    TALL,
		options: make(map[string]int32),
	}
}

func (b *Decaf) getDescription() string {
	return fmt.Sprintf("디카페인 커피 %s", b.size)
}

func (b *Decaf) cost() float32 {
	return 1.05
}

func (b *Decaf) setSize(size Size) {
	b.size = size
}

func (b *Decaf) getSize() Size {
	return b.size
}

func (b *Decaf) getOptions() map[string]int32 {
	return b.options
}

func (b *Decaf) setOptions(name string) {
	b.options[name] = b.options[name] + 1
}
