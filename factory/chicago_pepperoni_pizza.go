package factory

import "log"

type ChicagoStylePepperoniPizza struct {
}

func (p *ChicagoStylePepperoniPizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewChicagoStylePepperoniPizza() *ChicagoStylePepperoniPizza {
	return &ChicagoStylePepperoniPizza{}
}

func (p *ChicagoStylePepperoniPizza) prepare() {
	log.Println("ChicagoStylePepperoniPizza prepare")
}

func (p *ChicagoStylePepperoniPizza) bake() {
	log.Println("ChicagoStylePepperoniPizza bake")
}

func (p *ChicagoStylePepperoniPizza) cut() {
	log.Println("ChicagoStylePepperoniPizza cut")
}

func (p *ChicagoStylePepperoniPizza) box() {
	log.Println("ChicagoStylePepperoniPizza box")
}
