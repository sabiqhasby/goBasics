package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	message := make(chan int, 3) //menyimpan 3 data
	go func() {
		for {
			i := <-message
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		message <- i
	}
}
