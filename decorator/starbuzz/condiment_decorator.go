package starbuzz

type CondimentDecorator struct {
	beverage Beverage
}

func (d CondimentDecorator) getDescription() string {
	return d.getDescription()
}

func (d CondimentDecorator) setSize(size Size) {
	d.beverage.setSize(size)
}

func (d CondimentDecorator) getSize() Size {
	return d.beverage.getSize()
}

func (d CondimentDecorator) cost() float32 {
	return d.beverage.cost()
}

func (d CondimentDecorator) getOptions() map[string]int32 {
	return d.beverage.getOptions()
}

func (d CondimentDecorator) setOptions(name string) {
	d.beverage.setOptions(name)
}
