package command

import "testing"

func TestLightOnCommand_Execute(t *testing.T) {
	remote := NewSimpleRemoteControl()
	lightOn := NewLightOnCommand(newLight("방"))

	remote.setCommand(lightOn)
	remote.buttonWasPressed()
}
