package main

import "fmt"

func main() {
	numbers := []int{1, 2, 4, 5, 6}

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformed)

}

func transformNumbers(number *[]int, transform func(int) int) []int {
	dNumber := []int{}
	for _, val := range *number {
		dNumber = append(dNumber, transform(val))
	}

	return dNumber
}
