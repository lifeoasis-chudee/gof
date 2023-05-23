package command

import "fmt"

type Stereo interface {
	on()
	off()
	setCd()
	setDvd()
	setRadio()
	setVolume(v int)
}

type stereo struct {
	name   string
	volume int
}

func (s *stereo) on() {
	fmt.Println(s.name + " on")
}

func (s *stereo) off() {
	fmt.Println(s.name + " off")
}

func (s *stereo) setCd() {
	fmt.Println(s.name + " setCD")
}

func (s *stereo) setDvd() {
	fmt.Println(s.name + " setDvd")
}

func (s *stereo) setRadio() {
	fmt.Println(s.name + " setRadio")
}

func (s *stereo) setVolume(v int) {
	s.volume = v
	fmt.Println(s.name + " setVolume")
}

func newStereo(name string) *stereo {
	return &stereo{name: name, volume: 0}
}
