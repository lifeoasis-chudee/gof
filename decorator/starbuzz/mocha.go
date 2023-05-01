package starbuzz

import (
	"fmt"
)

type Mocha struct {
	beverage Beverage
}

func NewMocha(beverage Beverage) Beverage {
	return &Mocha{
		beverage: beverage,
	}
}

func (d *Mocha) getDescription() string {
	return fmt.Sprintf("%s, 모카", d.beverage.getDescription())
}

func (d *Mocha) cost() float32 {
	cost := d.beverage.cost()

	switch d.beverage.getSize() {
	case TALL:
		cost += 0.2
	case GRANDE:
		cost += 0.25
	case VENTI:
		cost += 0.3
	}

	return cost
}

func (d *Mocha) setSize(size Size) {
	d.beverage.setSize(size)
}

func (d *Mocha) getSize() Size {
	return d.beverage.getSize()
}
