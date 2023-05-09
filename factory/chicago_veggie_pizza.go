package factory

import "log"

type ChicagoStyleVeggiePizza struct {
}

func (p *ChicagoStyleVeggiePizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewChicagoStyleVeggiePizza() *ChicagoStyleVeggiePizza {
	return &ChicagoStyleVeggiePizza{}
}

func (p *ChicagoStyleVeggiePizza) prepare() {
	log.Println("ChicagoStyleVeggiePizza prepare")
}

func (p *ChicagoStyleVeggiePizza) bake() {
	log.Println("ChicagoStyleVeggiePizza bake")
}

func (p *ChicagoStyleVeggiePizza) cut() {
	log.Println("ChicagoStyleVeggiePizza cut")
}

func (p *ChicagoStyleVeggiePizza) box() {
	log.Println("ChicagoStyleVeggiePizza box")
}
