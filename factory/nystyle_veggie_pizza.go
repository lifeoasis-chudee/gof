package factory

import "log"

type NYStyleVeggiePizza struct {
}

func (p *NYStyleVeggiePizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewNYStyleVeggiePizza() *NYStyleVeggiePizza {
	return &NYStyleVeggiePizza{}
}

func (p *NYStyleVeggiePizza) prepare() {
	log.Println("NYStyleVeggiePizza prepare")
}

func (p *NYStyleVeggiePizza) bake() {
	log.Println("NYStyleVeggiePizza bake")
}

func (p *NYStyleVeggiePizza) cut() {
	log.Println("NYStyleVeggiePizza cut")
}

func (p *NYStyleVeggiePizza) box() {
	log.Println("NYStyleVeggiePizza box")
}
