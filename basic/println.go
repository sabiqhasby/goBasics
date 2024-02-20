package main

import "fmt"


func main(){

	var nameOne string = "Hello Sab"
	var nameTwo = "hasby"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Bowwer"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yogiii"

	fmt.Println(nameFour)


	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	var numOne int8 = 25
	var numTwo = -128
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)


	var scoreOne float32 = 28.78
	var scoreTwo float64 = 7889899.99

	fmt.Println(scoreOne, scoreTwo)
}