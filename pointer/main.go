package main

import "fmt"

func updateName(name *string) {
	*name = "Suteja"
	// return *name
}
func main() {
	var numberA int = 4

	var numberB = &numberA

	*numberB = 8

	fmt.Println("number a", numberA)
	fmt.Println("number b", *numberB)

	name := "tifa"

	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value is:", *m)

	fmt.Println("before update:", name)
	updateName(m)
	fmt.Println("after update: ", name)

}
