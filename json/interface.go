package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{"Name" : "John Doe", "Age":30}`
	var jsonData = []byte(jsonString)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)
	fmt.Println("User: ", data1["Name"])
	fmt.Println("Age: ", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	var decodedData = data2.(map[string]interface{})
	fmt.Println("User: ", decodedData["Name"])
	fmt.Println("Age: ", decodedData["Age"])

}
