package factory

import "log"

type cheesePizza struct {
}

func (p *cheesePizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewCheesePizza() *cheesePizza {
	return &cheesePizza{}
}

func (p *cheesePizza) prepare() {
	log.Println("cheesePizza prepare")
}

func (p *cheesePizza) bake() {
	log.Println("cheesePizza bake")
}

func (p *cheesePizza) cut() {
	log.Println("cheesePizza cut")
}

func (p *cheesePizza) box() {
	log.Println("cheesePizza box")
}
