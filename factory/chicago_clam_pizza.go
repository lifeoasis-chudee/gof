package factory

import "log"

type ChicagoStyleClamPizza struct {
}

func (p *ChicagoStyleClamPizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewChicagoStyleClamPizza() *ChicagoStyleClamPizza {
	return &ChicagoStyleClamPizza{}
}

func (p *ChicagoStyleClamPizza) prepare() {
	log.Println("ChicagoStyleClamPizza prepare")
}

func (p *ChicagoStyleClamPizza) bake() {
	log.Println("ChicagoStyleClamPizza bake")
}

func (p *ChicagoStyleClamPizza) cut() {
	log.Println("ChicagoStyleClamPizza cut")
}

func (p *ChicagoStyleClamPizza) box() {
	log.Println("ChicagoStyleClamPizza box")
}
