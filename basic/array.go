package main

import "fmt"

func main(){
	// var ages [3]int = [3]int{20,34, 39}
	var ages = [3]int{20,34, 39}
	names := [4]string{"Hasby", "Al", "Sabiq", "Naqi"}
	names[1] = "Luigi"
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(ages))
	

	var scores = []int{100, 50, 40}
	scores[2] = 25

	scores = append(scores, 75) //push to array
	fmt.Println("Array: ", scores)

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne,  "Mike")
	fmt.Println(rangeOne)
}