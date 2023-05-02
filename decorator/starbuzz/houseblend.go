package starbuzz

import "fmt"

type HouseBlend struct {
	size    Size
	options map[string]int32
}

func NewHouseBlend() Beverage {
	return &HouseBlend{
		size:    TALL,
		options: make(map[string]int32),
	}
}

func (b *HouseBlend) getDescription() string {
	return fmt.Sprintf("하우스블랜드 커피 %s", b.size)
}

func (b *HouseBlend) cost() float32 {
	return 0.89
}

func (b *HouseBlend) setSize(size Size) {
	b.size = size
}

func (b *HouseBlend) getSize() Size {
	return b.size
}

func (b *HouseBlend) getOptions() map[string]int32 {
	return b.options
}

func (b *HouseBlend) setOptions(name string) {
	b.options[name] = b.options[name] + 1
}
