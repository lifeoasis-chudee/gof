package factory

type NYStyleCheesePizza struct {
	*AbstractPizza
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	as := &AbstractPizza{
		name:     "뉴욕 스타일 소스와 치즈 피자",
		dough:    "씬 크러스트 도우",
		sauce:    "마리나라 소스",
		toppings: []string{"잘게 썬 레지아노 치즈"},
	}
	nys := &NYStyleCheesePizza{as}
	return nys
}
