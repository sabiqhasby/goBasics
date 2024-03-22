package main

import "fmt"

//recursion berarti function yg dipanggil oleh dirinya sendiri

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result

	// recursion!!!
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)

}
