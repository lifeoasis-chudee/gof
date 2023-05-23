package command

import (
	"testing"
)

func TestRemoteControl(t *testing.T) {
	remoteControl := NewRemoteControl()

	livingRoomLight := newLight("Living Room")
	kitchenLight := newLight("Kitchen")
	ceilingFan := newCeilingFan("Living Room")
	//garageDoor := newGarageDoor("Garage")
	stereo := newStereo("Living Room")

	livingRoomLightOn := NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := NewLightOffCommand(livingRoomLight)

	kitchenLightOn := NewLightOnCommand(kitchenLight)
	kitchenLightOff := NewLightOffCommand(kitchenLight)

	ceilingFanOn := NewCeilingFanOnCommand(ceilingFan)
	ceilingFanOff := NewCeilingFanOffCommand(ceilingFan)

	//garageDoorUp := NewGarageDoorUpCommand(garageDoor)
	//garageDoorDown := NewGarageDoorDownCommand(garageDoor)

	stereoOnWithCD := NewStereoOnWithCDCommand(stereo)
	stereoOff := NewStereoOffCommand(stereo)

	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.SetCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.SetCommand(2, ceilingFanOn, ceilingFanOff)
	remoteControl.SetCommand(3, stereoOnWithCD, stereoOff)

	remoteControl.String()

	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(1)
	remoteControl.OffButtonWasPushed(1)
	remoteControl.OnButtonWasPushed(2)
	remoteControl.OffButtonWasPushed(2)
	remoteControl.OnButtonWasPushed(3)
	remoteControl.OffButtonWasPushed(3)
}
