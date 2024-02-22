package main

import (
	"fmt"
	"runtime"
)

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, v := range numbers {
		sum += v
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	max := numbers[0]
	for _, v := range numbers {
		if max < v {
			fmt.Printf("max %v <  %v\n", max, v)
			max = v
		}
	}
	ch <- max
}

func main() {
	runtime.GOMAXPROCS(8)

	numbers := []int{3, 4, 4, 5, 6, 3, 2, 2, 6, 8, 3, 4, 6, 3}
	fmt.Println("numbers : ", numbers)

	ch1 := make(chan float64)
	go getAverage(numbers, ch1)

	ch2 := make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}

}
