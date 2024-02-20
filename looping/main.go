package main

import "fmt"

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value x:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value i:", i)
	// }

	names := []string{"yoshi", "bower", "sueb", "cici"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("now is at index: %v, and value is: %v \n", index, value)
	// }
	for _, value := range names {
		fmt.Printf("now value is: %v \n", value)
	}

}
