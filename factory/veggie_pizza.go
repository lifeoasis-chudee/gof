package factory

import "log"

type veggiePizza struct {
}

func (p *veggiePizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewVeggiePizza() *veggiePizza {
	return &veggiePizza{}
}

func (p *veggiePizza) prepare() {
	log.Println("veggiePizza prepare")
}

func (p *veggiePizza) bake() {
	log.Println("veggiePizza bake")
}

func (p *veggiePizza) cut() {
	log.Println("veggiePizza cut")
}

func (p *veggiePizza) box() {
	log.Println("veggiePizza box")
}
