package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := context.Background()

	ctx, cancel := context.WithTimeout(c, time.Second*2)
	defer cancel()
	fbreceiver := make(chan string)
	dbreceiver := make(chan string)

	go GetDataFromFb(ctx, fbreceiver)
	go GetDataFromDb(ctx, dbreceiver)

	defer close(fbreceiver)
	defer close(dbreceiver)

	for i := 0; i < 2; i++ {
		select {
		case fb := <-fbreceiver:
			fmt.Println(">>>>> Data received from: ", fb)
		case db := <-dbreceiver:
			fmt.Println(">>>>>> Data received from: ", db)
		case <-ctx.Done():
			fmt.Println(">>>>> TImeout to process")
		}
	}
	time.Sleep(time.Second * 50)
}

func GetDataFromDb(ctx context.Context, datareceiver chan string) {
	startTime := time.Now()
	ticker := time.NewTicker(time.Second * 1)

	for _ = range ticker.C {
		fmt.Println("still processing from DB")
		if time.Since(startTime).Seconds() > 10 {
			ticker.Stop()
			break
		}
	}
	if ctx.Err() == nil {
		datareceiver <- "database"
	}
	return

}

func GetDataFromFb(ctx context.Context, datareceiver chan string) {
	startTime := time.Now()
	ticker := time.NewTicker(time.Second * 1)

	for _ = range ticker.C {
		fmt.Println("still processing form FB")
		if time.Since(startTime).Seconds() >= 5 { // timeout after 5 seconds
			ticker.Stop()
			break
		}
	}
	if ctx.Err() == nil {
		datareceiver <- "facebook"
	}
	return
}
