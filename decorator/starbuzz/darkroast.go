package starbuzz

import "fmt"

type DarkRoast struct {
	size Size
}

func NewDarkRoast() Beverage {
	return &DarkRoast{}
}

func (d *DarkRoast) getDescription() string {
	return fmt.Sprintf("다크로스트 커피 %s", d.size)
}

func (d *DarkRoast) cost() float32 {
	return 0.99
}

func (d *DarkRoast) setSize(size Size) {
	d.size = size
}

func (d *DarkRoast) getSize() Size {
	return d.size
}
