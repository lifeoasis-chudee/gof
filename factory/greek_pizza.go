package factory

import "log"

type greekPizza struct {
}

func (p *greekPizza) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewGreekPizza() *greekPizza {
	return &greekPizza{}
}

func (p *greekPizza) prepare() {
	log.Println("greekPizza prepare")
}

func (p *greekPizza) bake() {
	log.Println("greekPizza bake")
}

func (p *greekPizza) cut() {
	log.Println("greekPizza cut")
}

func (p *greekPizza) box() {
	log.Println("greekPizza box")
}
