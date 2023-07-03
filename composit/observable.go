package composit

import (
	"fmt"
	"reflect"
)

type QuackObservable interface {
	registerObserver(observer Observer)
	notifyObservers()
}

type Observer interface {
	update(duck QuackObservable)
}

type Observable struct {
	observers []Observer
	duck      QuackObservable
	QuackObservable
}

func NewObservable(duck QuackObservable) *Observable {
	return &Observable{duck: duck}
}
func (o *Observable) registerObserver(observer Observer) {
	o.observers = append(o.observers, observer)
}
func (o *Observable) notifyObservers() {
	for i := range o.observers {
		o.observers[i].update(o.duck)
	}
}

type Quackollogist struct {
	Observer
}

func NewQuackollogist() *Quackollogist {
	return &Quackollogist{}
}
func (q *Quackollogist) update(duck QuackObservable) {
	fmt.Println(fmt.Sprintf("꽥꽥학자: %v가 방금 소리냈다.", reflect.TypeOf(duck)))
}
