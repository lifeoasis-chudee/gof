package starbuzz

type Decorator struct {
	beverage Beverage
}

func NewDecorator(beverage Beverage) *Decorator {
	return &Decorator{beverage}
}

func (d Decorator) getDescription() string {
	return d.beverage.getDescription()
}

func (d Decorator) cost() float32 {
	return d.beverage.cost()
}
