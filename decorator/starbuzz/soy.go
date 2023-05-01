package starbuzz

import "fmt"

type Soy struct {
	beverage Beverage
}

func NewSoy(beverage Beverage) Beverage {
	return &Soy{
		beverage: beverage,
	}
}

func (d *Soy) getDescription() string {
	return fmt.Sprintf("%s, 두유", d.beverage.getDescription())
}

func (d *Soy) cost() float32 {
	cost := d.beverage.cost()

	switch d.beverage.getSize() {
	case TALL:
		cost += 0.15
	case GRANDE:
		cost += 0.2
	case VENTI:
		cost += 0.25
	}

	return cost
}

func (d *Soy) setSize(size Size) {
	d.beverage.setSize(size)
}

func (d *Soy) getSize() Size {
	return d.beverage.getSize()
}
