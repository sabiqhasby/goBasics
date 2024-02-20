package main

import "fmt"

func main() {
	defer fmt.Println("halo") //defer akan di eksekusi terakhir setelah yg lain selesai
	fmt.Println("selamat datang")
}
