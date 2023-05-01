package starbuzz

import "fmt"

type HouseBlend struct {
	size Size
}

func NewHouseBlend() Beverage {
	return &HouseBlend{}
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
