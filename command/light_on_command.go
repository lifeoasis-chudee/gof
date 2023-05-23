package command

type LightOnCommand struct {
	Light
}

func NewLightOnCommand(light Light) *LightOnCommand {
	return &LightOnCommand{light}
}

func (l LightOnCommand) Execute() {
	l.on()
}

func (l LightOnCommand) Undo() {
	l.off()
}
