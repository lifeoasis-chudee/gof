package starbuzz

import "fmt"

type Decaf struct {
	size Size
}

func NewDecaf() Beverage {
	return &Decaf{}
}

func (d *Decaf) getDescription() string {
	return fmt.Sprintf("디카페인 커피 %s", d.size)
}

func (d *Decaf) cost() float32 {
	return 1.05
}

func (d *Decaf) setSize(size Size) {
	d.size = size
}

func (d *Decaf) getSize() Size {
	return d.size
}
