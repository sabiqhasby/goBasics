package main

import "fmt"

func main() {

	age := "35"
	name := "Hasby"
	fmt.Print("hello, ")
	fmt.Print("there \n")
	fmt.Print("hass, ")

	fmt.Println("Helloo, my Age is", age, "and, my Name is", name)

	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("my age is %T \n", age)
	fmt.Printf("my scored %f points! \n", 225.5)
	fmt.Printf("my scored %0.1f points! \n", 225.5)

	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}