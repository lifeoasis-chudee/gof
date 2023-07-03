package composit

var numberOfQuack = 0

type QuackCounter struct {
	duck Quackable
	Quackable
}

func NewQuackCounter(duck Quackable) *QuackCounter {
	return &QuackCounter{duck: duck}
}
func (qc *QuackCounter) quack() {
	qc.duck.quack()
	qc.notifyObservers()
	numberOfQuack++
}
func (qc *QuackCounter) getQuacks() int {
	return numberOfQuack
}
func (qc *QuackCounter) registerObserver(ob Observer) {
	qc.duck.registerObserver(ob)
}
func (qc *QuackCounter) notifyObservers() {}
