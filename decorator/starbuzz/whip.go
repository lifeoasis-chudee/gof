package starbuzz

import "fmt"

type Whip struct {
	beverage Beverage
}

func NewWhip(beverage Beverage) Beverage {
	return &Whip{
		beverage: beverage,
	}
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

func (d *Whip) setSize(size Size) {
	d.beverage.setSize(size)
}

func (d *Whip) getSize() Size {
	return d.beverage.getSize()
}
