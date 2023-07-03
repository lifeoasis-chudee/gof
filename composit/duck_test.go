package composit

import (
	"fmt"
	"testing"
)

func simulate(duck Quackable) {
	duck.quack()
}

func TestDuck(t *testing.T) {
	mallardDuck := NewMallardDuck()
	redDuck := NewRedheadDuck()
	duckCall := NewDuckCall()
	rubberDuck := NewRubberDuck()
	gooseDuck := NewGooseAdapter(Goose{})

	fmt.Println("\n오리 시뮬레이션 게임")

	quackologist := NewQuackollogist()
	mallardDuck.registerObserver(quackologist)

	simulate(mallardDuck)
	simulate(redDuck)
	simulate(duckCall)
	simulate(rubberDuck)
	simulate(gooseDuck)
}

func TestDuck_Counter(t *testing.T) {
	mallardDuck := NewQuackCounter(NewMallardDuck())
	redDuck := NewQuackCounter(NewRedheadDuck())
	duckCall := NewQuackCounter(NewDuckCall())
	rubberDuck := NewQuackCounter(NewRubberDuck())
	gooseDuck := NewGooseAdapter(Goose{})

	fmt.Println("\n오리 시뮬레이션 게임 (+데코레이터)")

	simulate(mallardDuck)
	simulate(redDuck)
	simulate(duckCall)
	simulate(rubberDuck)
	simulate(gooseDuck)

	fmt.Println(fmt.Sprintf("오리가 소리 낸 횟수: %d번", numberOfQuack))
}

func TestDuck_Factory(t *testing.T) {
	df := &AbstractDuckFactory{}

	mallardDuck := df.createMallardDuck()
	redDuck := df.createRedheadDuck()
	duckCall := df.createDuckCall()
	rubberDuck := df.createRubberDuck()
	gooseDuck := df.createGooseAdapter()

	fmt.Println("\n오리 시뮬레이션 게임 (+추상 팩토리)")

	simulate(mallardDuck)
	simulate(redDuck)
	simulate(duckCall)
	simulate(rubberDuck)
	simulate(gooseDuck)

	fmt.Println(fmt.Sprintf("오리가 소리 낸 횟수: %d번", numberOfQuack))
}

func TestDuck_Composit(t *testing.T) {
	df := &AbstractDuckFactory{}

	fmt.Println("\n오리 시뮬레이션 게임: 무리 (+컴포지트)")

	flockOfDucks := NewFlock()
	flockOfDucks.add(df.createRedheadDuck())
	flockOfDucks.add(df.createDuckCall())
	flockOfDucks.add(df.createRubberDuck())
	flockOfDucks.add(df.createGooseAdapter())

	flockOfMallards := NewFlock()
	flockOfMallards.add(df.createMallardDuck())
	flockOfMallards.add(df.createMallardDuck())
	flockOfMallards.add(df.createMallardDuck())
	flockOfMallards.add(df.createMallardDuck())

	flockOfDucks.add(flockOfMallards)

	fmt.Println("\n오리 시뮬레이션 게임: 전체 무리")
	simulate(flockOfDucks)

	fmt.Println("\n오리 시뮬레이션 게임: 물오리 무리")
	simulate(flockOfMallards)

	fmt.Println(fmt.Sprintf("오리가 소리 낸 횟수: %d번", numberOfQuack))
}

func TestDuck_Observer(t *testing.T) {
	df := &AbstractDuckFactory{}

	flockOfDucks := NewFlock()
	flockOfDucks.add(df.createRedheadDuck())
	flockOfDucks.add(df.createDuckCall())
	flockOfDucks.add(df.createRubberDuck())
	flockOfDucks.add(df.createGooseAdapter())

	flockOfMallards := NewFlock()
	flockOfMallards.add(df.createMallardDuck())
	flockOfMallards.add(df.createMallardDuck())
	flockOfMallards.add(df.createMallardDuck())
	flockOfMallards.add(df.createMallardDuck())

	flockOfDucks.add(flockOfMallards)

	fmt.Println("\n오리 시뮬레이션 게임: (+옵저버)")

	quackologist := NewQuackollogist()
	flockOfDucks.registerObserver(quackologist)

	simulate(flockOfDucks)

	fmt.Println(fmt.Sprintf("오리가 소리 낸 횟수: %d번", numberOfQuack))
}
