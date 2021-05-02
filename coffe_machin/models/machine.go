package models

type Machine struct {
	Water int
	Milk int
	CoffeeBeans int
	Cups int
	Money int

	CoffeTypes map[string]Coffe
}

func (m *Machine) Fill(plusWater, plusMilk, plusBeans, plusCups int) {
	m.Water += plusWater
	m.Milk += plusMilk
	m.CoffeeBeans += plusBeans
	m.Cups += plusCups
}
func (m * Machine) UnFill(minusWater, minusMilk, minusBeans, minusCups, plusMoney int) {
	m.Water -= minusWater
	m.Milk -= minusMilk
	m.CoffeeBeans -= minusBeans
	m.Cups -= minusCups
	m.Money += plusMoney

}