package main

import "fmt"

func main() {
	var numberA int = 4

	var numberB = &numberA

	*numberB = 8

	fmt.Println("number a", numberA)
	fmt.Println("number b", *numberB)
}
