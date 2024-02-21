package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string
	for _, val := range names {
		initials = append(initials, val[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("John doe")
	fmt.Println(fn1, sn1)
	fn2, sn2 := getInitials("Cloud strike")
	fmt.Println(fn2, sn2)
	fn3, sn3 := getInitials("strike")
	fmt.Println(fn3, sn3)
}
