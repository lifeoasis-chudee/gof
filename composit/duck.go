package composit

import (
	"fmt"
)

type Quackable interface {
	quack()
	QuackObservable
}

type MallardDuck struct {
	observable *Observable
	Quackable
}

func NewMallardDuck() *MallardDuck {
	md := &MallardDuck{}
	md.observable = NewObservable(md)
	return md
}
func (d *MallardDuck) quack() {
	fmt.Println("꽥꽥")
	d.notifyObservers()
}
func (d *MallardDuck) registerObserver(ob Observer) {
	d.observable.registerObserver(ob)
}
func (d *MallardDuck) notifyObservers() {
	d.observable.notifyObservers()
}

type RedheadDuck struct {
	observable *Observable
	Quackable
}

func NewRedheadDuck() *RedheadDuck {
	rd := &RedheadDuck{}
	rd.observable = NewObservable(rd)
	return rd
}
func (d *RedheadDuck) quack() {
	fmt.Println("꽥꽥")
	d.notifyObservers()
}
func (d *RedheadDuck) registerObserver(ob Observer) {
	d.observable.registerObserver(ob)
}
func (d *RedheadDuck) notifyObservers() {
	d.observable.notifyObservers()
}

type DuckCall struct {
	observable *Observable
	Quackable
}

func NewDuckCall() *DuckCall {
	dc := &DuckCall{}
	dc.observable = NewObservable(dc)
	return dc
}
func (d *DuckCall) quack() {
	fmt.Println("꽉꽉")
	d.notifyObservers()
}
func (d *DuckCall) registerObserver(ob Observer) {
	d.observable.registerObserver(ob)
}
func (d *DuckCall) notifyObservers() {
	d.observable.notifyObservers()
}

type RubberDuck struct {
	observable *Observable
	Quackable
}

func NewRubberDuck() *RubberDuck {
	rd := &RubberDuck{}
	rd.observable = NewObservable(rd)
	return rd
}
func (d *RubberDuck) quack() {
	fmt.Println("삑삑")
	d.notifyObservers()
}
func (d *RubberDuck) registerObserver(ob Observer) {
	d.observable.registerObserver(ob)
}
func (d *RubberDuck) notifyObservers() {
	d.observable.notifyObservers()
}

type Goose struct{}

func (g Goose) honk() {
	fmt.Println("끽끽")
}

type GooseAdapter struct {
	goose      Goose
	observable *Observable
	Quackable
}

func NewGooseAdapter(goose Goose) Quackable {
	ga := &GooseAdapter{goose: goose}
	ga.observable = NewObservable(ga)
	return ga
}

func (a *GooseAdapter) quack() {
	a.goose.honk()
	a.notifyObservers()
}
func (a *GooseAdapter) registerObserver(ob Observer) {
	a.observable.registerObserver(ob)
}
func (a *GooseAdapter) notifyObservers() {
	a.observable.notifyObservers()
}
