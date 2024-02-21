package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill breakdown: \n"

	var total float64 = 0
	for item, price := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", item+":", price)
		total += price
	}

	//addtip
	fs += fmt.Sprintf("%-25v ...$%.2f \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%.2f \n", "total:", total+b.tip)

	return fs
}

//update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip

}

func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}
