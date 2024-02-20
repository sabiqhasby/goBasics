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
	var jsonString = `[{"Name" : "John Doe", "Age":30}, {"Name": "Simon", "Age": 44}]`
	var jsonData = []byte(jsonString)

	var data []User
	json.Unmarshal(jsonData, &data)
	fmt.Println("User 1: ", data[0].FullName)
	fmt.Println("User 2:", data[1].FullName)

}
