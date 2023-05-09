package factory

type SimplePizzaFactory struct {
}

func (f *SimplePizzaFactory) CreatePizza(kind string) Pizza {
	var pizza Pizza

	switch kind {
	case "cheese":
		pizza = NewCheesePizza()
	case "greek":
		pizza = NewGreekPizza()
	case "pepperoni":
		pizza = NewPepperoniPizza()
	case "clam":
		pizza = NewClamPizza()
	case "veggie":
		pizza = NewVeggiePizza()
	}

	return pizza
}
