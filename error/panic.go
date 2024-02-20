package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Input cannot be empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("type your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		defer fmt.Println("tetap terbaca") // masih di baca oleh program
		panic(err.Error())
		fmt.Println("end") // tidak akan terbaca

	}

}
