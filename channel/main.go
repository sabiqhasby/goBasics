package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var message = make(chan string)
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		message <- data //memasukkan data ke dalam channel
	}

	go sayHelloTo("John wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("json statham")

	var message1 = <-message //memasukkan isi message ke dalam variabel baru
	fmt.Println(message1)
	var message2 = <-message //memasukkan isi message ke dalam variabel baru
	fmt.Println(message2)
	var message3 = <-message //memasukkan isi message ke dalam variabel baru
	fmt.Println(message3)
}
