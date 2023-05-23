package command

import (
	"fmt"
	"testing"
)

func TestNewRemoteLambdaLoader(t *testing.T) {
	loader := NewRemoteLambdaLoader()

	loader.remoteControl.OnButtonWasPushed(0)
	loader.remoteControl.OffButtonWasPushed(0)
}

func TestNewRemoteControl(t *testing.T) {
	loader := NewRemoteLoader()

	loader.remoteControl.OnButtonWasPushed(0)
	loader.remoteControl.OffButtonWasPushed(0)

	loader.remoteControl.String()
	loader.remoteControl.UndoButtonWasPushed()
	fmt.Println()

	loader.remoteControl.OffButtonWasPushed(0)
	loader.remoteControl.OnButtonWasPushed(0)

	loader.remoteControl.String()
	loader.remoteControl.UndoButtonWasPushed()
}
