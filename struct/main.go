package main

import "fmt"

type student struct {
	name  string
	grade int
}

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

func main() {
	// var s1 student
	s1 := student{}
	s1.name = "my skill"
	s1.grade = 2

	fmt.Println("name", s1.name)
	fmt.Println("grade", s1.grade)

	b1 := newBill("John")
	fmt.Println(b1)
}
