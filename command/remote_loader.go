package command

type RemoteLoader struct {
	remoteControl *RemoteControl
}

func NewRemoteLambdaLoader() *RemoteLoader {
	remoteControl := NewRemoteControl()

	livingRoomLight := newLight("Living Room")

	remoteControl.SetCommand(0,
		CommandFunc(livingRoomLight.on),
		CommandFunc(livingRoomLight.off),
	)

	return &RemoteLoader{remoteControl: remoteControl}
}

func NewRemoteLoader() *RemoteLoader {
	remoteControl := NewRemoteControl()

	livingRoomLight := newLight("Living Room")

	livingRoomLightOn := NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := NewLightOffCommand(livingRoomLight)

	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)

	return &RemoteLoader{remoteControl: remoteControl}
}
