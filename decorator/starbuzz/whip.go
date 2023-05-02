package starbuzz

import "fmt"

type Whip struct {
	CondimentDecorator
}

func NewWhip(beverage Beverage) Beverage {
	b := CondimentDecorator{beverage}
	b.setOptions("whip")
	return &Whip{b}
}

func (d *Whip) getDescription() string {
	return fmt.Sprintf("%s, 휘핑크림", d.beverage.getDescription())
}

func (d *Whip) cost() float32 {
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
