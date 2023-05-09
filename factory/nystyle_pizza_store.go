package factory

type NYStylePizzaStore struct {
	*AbstractPizzaStore
}

func NewNYStylePizzaStore() *NYStylePizzaStore {
	as := &AbstractPizzaStore{}
	r := &NYStylePizzaStore{as}
	r.AbstractPizzaStore.CreatePizza = r.CreatePizza
	return r
}

func (s *NYStylePizzaStore) CreatePizza(kind string) Pizza {
	var pizza Pizza

	switch kind {
	case "cheese":
		pizza = NewNYStyleCheesePizza()
	case "pepperoni":
		pizza = NewNYStylePepperoniPizza()
	case "clam":
		pizza = NewNYStyleClamPizza()
	case "veggie":
		pizza = NewNYStyleVeggiePizza()
	}

	return pizza
}
