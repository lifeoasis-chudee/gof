package starbuzz

type Size string

var (
	TALL   Size = "tall"
	GRANDE Size = "grande"
	VENTI  Size = "venti"
)

type Beverage interface {
	getDescription() string
	cost() float32
	setSize(size Size)
	getSize() Size
}
