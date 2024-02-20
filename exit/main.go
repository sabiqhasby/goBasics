package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("halo")
	os.Exit(1) //memberhentikan program secara paksa
	fmt.Println("selamat datang")
}
