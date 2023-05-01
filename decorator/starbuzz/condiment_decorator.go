package starbuzz

type CondimentDecorator interface {
	Beverage
	getDescription() string
	getSize() Size
}
