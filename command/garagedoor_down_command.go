package command

type GarageDoorDownCommand struct {
	GarageDoor
}

func NewGarageDoorDownCommand(garageDoor GarageDoor) *GarageDoorDownCommand {
	return &GarageDoorDownCommand{garageDoor}
}

func (g GarageDoorDownCommand) Execute() {
	g.down()
}

func (g GarageDoorDownCommand) Undo() {
	g.up()
}
