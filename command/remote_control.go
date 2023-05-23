package command

import (
	"fmt"
	"reflect"
	"strings"
)

type RemoteControl struct {
	onCommands  []Command
	offCommands []Command
	undoCommand Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		onCommands:  make([]Command, 7),
		offCommands: make([]Command, 7),
		undoCommand: nil,
	}
}

func (r *RemoteControl) SetCommand(slot int, onCommand Command, offCommand Command) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r *RemoteControl) OnButtonWasPushed(slot int) {
	if r.onCommands[slot] != nil {
		r.onCommands[slot].Execute()
		r.undoCommand = r.onCommands[slot]
	}
}

func (r *RemoteControl) OffButtonWasPushed(slot int) {
	if r.offCommands[slot] != nil {
		r.offCommands[slot].Execute()
		r.undoCommand = r.offCommands[slot]
	}
}

func (r *RemoteControl) UndoButtonWasPushed() {
	if r.undoCommand != nil {
		r.undoCommand.Undo()
	}
}

func (r *RemoteControl) String() {
	fmt.Print("\n------ 리모컨 ------\n")
	for i, onCommand := range r.onCommands {
		onCommandType := "NoCommand"
		if reflect.TypeOf(onCommand) != nil {
			onCommandType = strings.TrimPrefix(reflect.TypeOf(onCommand).String(), "*command.")
		}

		offCommandType := "NoCommand"
		if reflect.TypeOf(r.offCommands[i]) != nil {
			offCommandType = strings.TrimPrefix(reflect.TypeOf(r.offCommands[i]).String(), "*command.")
		}

		fmt.Printf("[slot %d] %-20s \t %-20s\n", i, onCommandType, offCommandType)
	}

	undoCommandType := "NoCommand"
	if reflect.TypeOf(r.undoCommand) != nil {
		undoCommandType = strings.TrimPrefix(reflect.TypeOf(r.undoCommand).String(), "*command.")
	}
	fmt.Printf("[undo] %-20s\n\n", undoCommandType)
}
