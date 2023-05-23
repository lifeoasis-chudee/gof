package command

type CeilingFanOffCommand struct {
	CeilingFan
}

func NewCeilingFanOffCommand(fan CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{fan}
}

func (l CeilingFanOffCommand) Execute() {
	l.off()
}

func (l CeilingFanOffCommand) Undo() {
	l.on()
}
