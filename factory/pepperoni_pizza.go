package factory

import "log"

type pepperoniPizza struct {
}

func (p *pepperoniPizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewPepperoniPizza() *pepperoniPizza {
	return &pepperoniPizza{}
}

func (p *pepperoniPizza) prepare() {
	log.Println("pepperoniPizza prepare")
}

func (p *pepperoniPizza) bake() {
	log.Println("pepperoniPizza bake")
}

func (p *pepperoniPizza) cut() {
	log.Println("pepperoniPizza cut")
}

func (p *pepperoniPizza) box() {
	log.Println("pepperoniPizza box")
}
