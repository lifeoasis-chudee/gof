package command

type StereoOnWithCDCommand struct {
	Stereo
}

func NewStereoOnWithCDCommand(stereo Stereo) *StereoOnWithCDCommand {
	return &StereoOnWithCDCommand{stereo}
}

func (l StereoOnWithCDCommand) Execute() {
	l.on()
	l.setCd()
	l.setVolume(11)
}

func (l StereoOnWithCDCommand) Undo() {
	l.off()
}
