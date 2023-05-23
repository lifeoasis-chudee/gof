package command

type GarageDoorUpCommand struct {
	GarageDoor
}

func NewGarageDoorUpCommand(garageDoor GarageDoor) *GarageDoorUpCommand {
	return &GarageDoorUpCommand{garageDoor}
}

func (g GarageDoorUpCommand) Execute() {
	g.up()
}

func (g GarageDoorUpCommand) Undo() {
	g.down()
}
