package main

import "fmt"

func main() {

	// map[key]value <<<<<

	// var chicken map[string]int
	chicken := map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])

	menu := map[string]float64{
		"soup":    4.99,
		"pie":     7.99,
		"salad":   6.99,
		"pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//Looping maps
	for key, val := range menu {
		fmt.Println(key, "-", val)
	}

	//ints as key type
	phonebook := map[int]string{
		2673333999: "mario",
		5443455566: "luigi",
		6655334234: " peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[2673333999]) // Mario

	//UPdate the map
	phonebook[2673333999] = "bowser"
	fmt.Println(phonebook)
}
