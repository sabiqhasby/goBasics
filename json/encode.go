package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var object = []User{{"John Doe", 30}, {"Simon", 44}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var data = string(jsonData)
	fmt.Println(data)

}
