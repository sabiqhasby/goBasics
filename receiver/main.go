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
		items: map[string]float64{"susu": 5.4, "pie": 5.3},
		tip:   8,
	}
	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"

	var total float64 = 0
	for item, price := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", item+":", price)
		total += price
	}

	//total
	fs += fmt.Sprintf("%-25v ...$%.2f", "total:", total)

	return fs
}

func main() {
	myBill := newBill("John Doe")
	fmt.Println(myBill.format())
}
