package factory

import "log"

type ChicagoStyleCheesePizza struct {
	*AbstractPizza
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	as := &AbstractPizza{
		name:     "시카고 스타이 딥 디쉬 치즈 피자",
		dough:    "아주 두꺼운 크러스트 도우",
		sauce:    "플럼토마토 소스",
		toppings: []string{"잘게 조각낸 모짜렐라 치즈"},
	}
	nys := &ChicagoStyleCheesePizza{as}
	return nys
}

func (p *ChicagoStyleCheesePizza) cut() {
	log.Println("네모난 모양으로 피자 가르기")
}
