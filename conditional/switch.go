package main

import "fmt"

func main() {
	var point = 5
	switch point {
	case 8, 9:
		fmt.Println("Perfect")
	case 7, 6, 5:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}
