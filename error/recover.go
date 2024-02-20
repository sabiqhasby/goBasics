// menghandle panic error
// saat panic error muncul, recover ini kan men-take-over goroutine yg sedang panic
// sehingga pesan panic tidak akan muncul

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

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("application running perfectly")
	}
}

func main() {
	defer catch()
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("halo", name)

	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
