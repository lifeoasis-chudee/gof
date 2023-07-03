package composit

type Flock struct {
	quackers []Quackable
	Quackable
}

func NewFlock() *Flock {
	return &Flock{}
}
func (f *Flock) add(quacker Quackable) {
	f.quackers = append(f.quackers, quacker)
}

func (f *Flock) quack() {
	for i := range f.quackers {
		f.quackers[i].quack()
	}
}
func (f *Flock) registerObserver(ob Observer) {
	for i := range f.quackers {
		f.quackers[i].registerObserver(ob)
	}
}
func (f *Flock) notifyObservers() {}
