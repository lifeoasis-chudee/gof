package factory

import "log"

type NYStyleClamPizza struct {
}

func (p *NYStyleClamPizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewNYStyleClamPizza() *NYStyleClamPizza {
	return &NYStyleClamPizza{}
}

func (p *NYStyleClamPizza) prepare() {
	log.Println("NYStyleClamPizza prepare")
}

func (p *NYStyleClamPizza) bake() {
	log.Println("NYStyleClamPizza bake")
}

func (p *NYStyleClamPizza) cut() {
	log.Println("NYStyleClamPizza cut")
}

func (p *NYStyleClamPizza) box() {
	log.Println("NYStyleClamPizza box")
}
