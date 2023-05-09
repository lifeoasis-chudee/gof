package main

import (
	"gof/factory"
	"log"
)

func main() {
	// 1
	log.Println("[OrderPizza]")
	factory.OrderPizza("cheese")
	factory.OrderPizza("veggie")
	factory.OrderPizza("pepperoni")
	factory.OrderPizza("clam")

	// 2
	log.Println("[SimplePizzaFactory]")
	ss := factory.SimplePizzaFactory{}
	ss.CreatePizza("cheese")
	ss.CreatePizza("veggie")
	ss.CreatePizza("pepperoni")
	ss.CreatePizza("clam")
	log.Println()

	// 3.1
	log.Println("[NYStylePizzaStore]")
	nys := factory.NewNYStylePizzaStore()
	p1 := nys.OrderPizza("cheese")
	log.Println("에단이 주문한", p1.Name())
	nys.OrderPizza("veggie")
	nys.OrderPizza("pepperoni")
	nys.OrderPizza("clam")
	log.Println()

	// 3.2
	log.Println("[ChicagoStylePizzaStore]")
	cs := factory.NewChicagoStylePizzaStore()
	p2 := cs.OrderPizza("cheese")
	log.Println("조엘이 주문한", p2.Name())
	cs.OrderPizza("veggie")
	cs.OrderPizza("pepperoni")
	cs.OrderPizza("clam")
	log.Println()
}
