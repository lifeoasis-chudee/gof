package command

import "testing"

func TestGarageDoorOpenCommand_Execute(t *testing.T) {
	remote := NewSimpleRemoteControl()
	light := newLight("방")
	garageDoor := newGarageDoor("Garage")
	lightOn := NewLightOnCommand(light)
	garageOpen := NewGarageDoorOpenCommand(garageDoor)

	remote.setCommand(lightOn)
	remote.buttonWasPressed()
	remote.setCommand(garageOpen)
	remote.buttonWasPressed()
}
