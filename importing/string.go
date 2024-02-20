package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "fri"))
	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("original:", greeting)

	ages := []int{45, 20, 30, 75, 24, 60, 65}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)

	fmt.Println(index) // returns the index of where to insert 30 in sorted ages slice

	names := []string{"yoshi", "bower", "sueb", "cici"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "sueb"))
}
