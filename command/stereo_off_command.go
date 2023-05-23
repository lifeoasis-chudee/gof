package command

type StereoOffCommand struct {
	Stereo
}

func NewStereoOffCommand(stereo Stereo) *StereoOffCommand {
	return &StereoOffCommand{stereo}
}

func (l StereoOffCommand) Execute() {
	l.off()
}

func (l StereoOffCommand) Undo() {
	l.on()
	l.setCd()
	l.setVolume(11)
}
