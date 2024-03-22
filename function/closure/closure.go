package main

import "fmt"

func main() {

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 2, 4, 5}

	var min, max = getMinMax(numbers)
	fmt.Println(min, max) // Output: 2 5

}
