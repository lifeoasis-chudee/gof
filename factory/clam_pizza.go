package factory

import "log"

type clamPizza struct {
}

func (p *clamPizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewClamPizza() *clamPizza {
	return &clamPizza{}
}

func (p *clamPizza) prepare() {
	log.Println("clamPizza prepare")
}

func (p *clamPizza) bake() {
	log.Println("clamPizza bake")
}

func (p *clamPizza) cut() {
	log.Println("clamPizza cut")
}

func (p *clamPizza) box() {
	log.Println("clamPizza box")
}
