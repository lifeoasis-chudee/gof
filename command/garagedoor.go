package command

import "fmt"

type GarageDoor interface {
	up()
	down()
	stop()
	lightOn()
	lightOff()
}

type garageDoor struct {
	name string
}

func newGarageDoor(name string) *garageDoor {
	return &garageDoor{name}
}

func (g garageDoor) up() {
	fmt.Println(g.name + " up")
}

func (g garageDoor) down() {
	fmt.Println(g.name + " down")
}

func (g garageDoor) stop() {
	fmt.Println(g.name + " stp[")
}

func (g garageDoor) lightOn() {
	fmt.Println(g.name + " lightOn")
}

func (g garageDoor) lightOff() {
	fmt.Println(g.name + " lightOff")
}
