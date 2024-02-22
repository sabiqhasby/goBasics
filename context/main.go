package main

import (
	"context"
	"fmt"
)

func main() {
	contextParent := context.Background()

	ctx1 := context.WithValue(contextParent, "key1", "hello world")
	ctx2 := context.WithValue(ctx1, "key2", "ngambil dari ctx1")
	ctx3 := context.WithValue(contextParent, "key3", "ngambil dari parent")

	fmt.Println(ctx1.Value("key1"))
	fmt.Println(ctx2.Value("key2"))
	fmt.Println(ctx3.Value("key1"))
}
