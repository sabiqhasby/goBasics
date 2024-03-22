package main

import "fmt"

func main() {
	numbers := []int{1, 2, 4, 5, 6}

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(number *[]int, transform func(int) int) []int {
	dNumber := []int{}
	for _, val := range *number {
		dNumber = append(dNumber, transform(val))
	}

	return dNumber
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
