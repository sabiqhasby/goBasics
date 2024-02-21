package main

import "fmt"

func updateName(x string) string {
	x = "Sutejo"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.49
}
func main() {
	//Grup A -> string, int, bool, float, array, struct
	name := "tifa"

	name = updateName(name)

	fmt.Println(name)

	//Grup B -> slices, maps, function
	//this group will not replacing the new value
	//but adding it to the list

	menu := map[string]float64{
		"pie":       7.9,
		"ice cream": 8.9,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
