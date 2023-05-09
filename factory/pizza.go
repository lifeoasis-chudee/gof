package factory

import "log"

type Pizza interface {
	prepare()
	bake()
	cut()
	box()
	Name() string
}

type AbstractPizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (p *AbstractPizza) prepare() {
	log.Println("준비 중:", p.name)
	log.Println("도우를 돌리는 중...")
	log.Println("소스를 뿌리는 중...")
	log.Println("토핑을 올리는 중...")
	for _, tp := range p.toppings {
		log.Println(" ", tp)
	}
}

func (p *AbstractPizza) bake() {
	log.Println("175도에서 25분 간 굽기")
}

func (p *AbstractPizza) cut() {
	log.Println("피자를 사선으로 자르기")
}

func (p *AbstractPizza) box() {
	log.Println("상자에 피자 담기")
}

func (p *AbstractPizza) Name() string {
	return p.name
}

func OrderPizza(kind string) Pizza {
	var pizza Pizza

	switch kind {
	case "cheese":
		pizza = NewCheesePizza()
	//case "greek":
	//	pizza = NewGreekPizza()
	case "pepperoni":
		pizza = NewPepperoniPizza()
	case "clam":
		pizza = NewClamPizza()
	case "veggie":
		pizza = NewVeggiePizza()
	}

	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza
}
