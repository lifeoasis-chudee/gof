package command

type LightOffCommand struct {
	Light
}

func NewLightOffCommand(light Light) *LightOffCommand {
	return &LightOffCommand{light}
}

func (l LightOffCommand) Execute() {
	l.off()
}

func (l LightOffCommand) Undo() {
	l.on()
}
