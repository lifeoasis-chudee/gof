package factory

type ChicagoStylePizzaStore struct {
	*AbstractPizzaStore
}

func NewChicagoStylePizzaStore() *ChicagoStylePizzaStore {
	as := &AbstractPizzaStore{}
	r := &ChicagoStylePizzaStore{as}
	r.AbstractPizzaStore.CreatePizza = r.CreatePizza
	return r
}

func (s *ChicagoStylePizzaStore) CreatePizza(kind string) Pizza {
	var pizza Pizza

	switch kind {
	case "cheese":
		pizza = NewChicagoStyleCheesePizza()
	case "pepperoni":
		pizza = NewChicagoStylePepperoniPizza()
	case "clam":
		pizza = NewChicagoStyleClamPizza()
	case "veggie":
		pizza = NewChicagoStyleVeggiePizza()
	}

	return pizza
}
