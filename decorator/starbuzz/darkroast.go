package starbuzz

import "fmt"

type DarkRoast struct {
	size    Size
	options map[string]int32
}

func NewDarkRoast() Beverage {
	return &DarkRoast{
		size:    TALL,
		options: make(map[string]int32),
	}
}

func (b *DarkRoast) getDescription() string {
	return fmt.Sprintf("다크로스트 커피 %s", b.size)
}

func (b *DarkRoast) cost() float32 {
	return 0.99
}

func (b *DarkRoast) setSize(size Size) {
	b.size = size
}

func (b *DarkRoast) getSize() Size {
	return b.size
}

func (b *DarkRoast) getOptions() map[string]int32 {
	return b.options
}

func (b *DarkRoast) setOptions(name string) {
	b.options[name] = b.options[name] + 1
}
