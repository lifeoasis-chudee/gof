package factory

type PizzaStore interface {
	CreatePizza(kind string) Pizza
	OrderPizza(kind string) Pizza
}

type AbstractPizzaStore struct {
	CreatePizza func(kind string) Pizza
}

func (s *AbstractPizzaStore) OrderPizza(kind string) Pizza {
	pizza := s.CreatePizza(kind)

	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza
}
