package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 4, 5, 6}
	moreNumb := []int{3, 5, 7, 8, 9}

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)

	gettransform1 := getTransformerFunction(&numbers)
	gettransform2 := getTransformerFunction(&moreNumb)

	transform1 := transformNumbers(&numbers, gettransform1)
	transform2 := transformNumbers(&moreNumb, gettransform2)

	fmt.Println("tr1:", transform1)
	fmt.Println("tr2:", transform2)

}

func transformNumbers(number *[]int, transform transformFn) []int {
	dNumber := []int{}
	for _, val := range *number {
		dNumber = append(dNumber, transform(val))
	}

	return dNumber
}

func getTransformerFunction(number *[]int) transformFn {
	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
