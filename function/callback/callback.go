package main

import "fmt"

func main() {
	var hasil = filter([]string{"ini", "data"}, func(string) bool {
		return true
	})
	fmt.Println(hasil)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, v := range data {

		if filtered := callback(v); filtered {
			result = append(result, v)
		}

	}
	return result
}
