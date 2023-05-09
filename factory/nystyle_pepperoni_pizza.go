package factory

import "log"

type NYStylePepperoniPizza struct {
}

func (p *NYStylePepperoniPizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewNYStylePepperoniPizza() *NYStylePepperoniPizza {
	return &NYStylePepperoniPizza{}
}

func (p *NYStylePepperoniPizza) prepare() {
	log.Println("NYStylePepperoniPizza prepare")
}

func (p *NYStylePepperoniPizza) bake() {
	log.Println("NYStylePepperoniPizza bake")
}

func (p *NYStylePepperoniPizza) cut() {
	log.Println("NYStylePepperoniPizza cut")
}

func (p *NYStylePepperoniPizza) box() {
	log.Println("NYStylePepperoniPizza box")
}
