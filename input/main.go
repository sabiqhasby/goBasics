package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("choose option (a - add item, s - save bill, t - add tip) ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("item Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be number")

			//back to input
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("enter tip amount: ", reader)
		fmt.Println(tip)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be number")

			//back to input
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("saved bill", b.name)

	default:
		fmt.Println("that not valid option")
		//back to input
		promptOptions(b)
	}

}
func main() {
	myBill := createBill()
	promptOptions(myBill)
	fmt.Println(myBill)
}
