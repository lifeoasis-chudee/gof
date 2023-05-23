package command

import "fmt"

type CeilingFan interface {
	on()
	off()
}

type ceilingFan struct {
	name string
}

func (s *ceilingFan) on() {
	fmt.Println(s.name + " on")
}

func (s *ceilingFan) off() {
	fmt.Println(s.name + " off")
}

func newCeilingFan(name string) *ceilingFan {
	return &ceilingFan{name: name}
}
