package starbuzz

import "fmt"

type Milk struct {
	CondimentDecorator
}

func NewMilk(beverage Beverage) Beverage {
	b := CondimentDecorator{beverage}
	b.setOptions("milk")
	return &Milk{b}
}

func (d *Milk) getDescription() string {
	return fmt.Sprintf("%s, 우유", d.beverage.getDescription())
}

func (d *Milk) cost() float32 {
	cost := d.beverage.cost()

	switch d.beverage.getSize() {
	case TALL:
		cost += 0.1
	case GRANDE:
		cost += 0.15
	case VENTI:
		cost += 0.2
	}

	return cost
}
