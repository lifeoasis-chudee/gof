package command

type CeilingFanOnCommand struct {
	CeilingFan
}

func NewCeilingFanOnCommand(fan CeilingFan) *CeilingFanOnCommand {
	return &CeilingFanOnCommand{fan}
}

func (l CeilingFanOnCommand) Execute() {
	l.on()
}

func (l CeilingFanOnCommand) Undo() {
	l.off()
}
