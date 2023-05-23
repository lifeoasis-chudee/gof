package command

type SimpleRemoteControl struct {
	slot Command
}

func NewSimpleRemoteControl() *SimpleRemoteControl {
	return &SimpleRemoteControl{}
}

func (c *SimpleRemoteControl) setCommand(command Command) {
	c.slot = command
}

func (c *SimpleRemoteControl) buttonWasPressed() {
	c.slot.Execute()
}
