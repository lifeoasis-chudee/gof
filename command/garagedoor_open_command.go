package command

type GarageDoorOpenCommand struct {
	GarageDoor
}

func NewGarageDoorOpenCommand(garageDoor GarageDoor) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{garageDoor}
}

func (g GarageDoorOpenCommand) Execute() {
	g.up()
	g.lightOn()
}

func (g GarageDoorOpenCommand) Undo() {
	g.stop()
	g.lightOff()
	g.down()
}
