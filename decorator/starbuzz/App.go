package starbuzz

import "log"

func App() {
	beverage := NewEspresso()
	beverage.setSize(TALL)
	log.Println(beverage.getDescription())

	beverage2 := NewDarkRoast()
	beverage2.setSize(VENTI)
	beverage2 = NewMocha(beverage2)
	beverage2 = NewMocha(beverage2)
	beverage2 = NewWhip(beverage2)
	log.Println(beverage2.getDescription(), " $", beverage2.cost())

	// "두유를 추가하고 휘핑크림을 얹은 더블 모카 한 잔 주세요!
	beverage3 := NewHouseBlend()
	beverage3.setSize(GRANDE)
	beverage3 = NewSoy(beverage3)
	beverage3 = NewMocha(beverage3)
	beverage3 = NewMocha(beverage3)
	beverage3 = NewWhip(beverage3)
	log.Println(beverage3.getDescription(), " $", beverage3.cost())
}
