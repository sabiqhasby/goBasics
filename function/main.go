package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Good Bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	// sayGreeting("mario")
	// sayGreeting("luigi")
	// sayBye("mario")

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloudgv", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.1f and cicle 2 is %0.3f", a1, a2)
}
