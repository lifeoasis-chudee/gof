package command

import (
	"fmt"
)

type Light interface {
	on()
	off()
}

type light struct {
	name string
}

func newLight(name string) *light {
	return &light{name: name}
}

func (l light) on() {
	fmt.Println(l.name + "의 조명이 켜졌습니다.")
}
func (l light) off() {
	fmt.Println(l.name + "의 조명이 꺼졌습니다.")
}
