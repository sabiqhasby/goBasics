package main

import (
	"fmt"
	"math"
)

func main() {
	area, circum := calculate(3.5)
	fmt.Println("area", area)
	fmt.Println("circum", circum)
}
func calculate(d float64) (float64, float64) {

	var area = math.Pi * math.Pow(d/2, 2) //luas
	var circumference = math.Pi * d       //keliling

	//kembalikan 2 nilai
	return area, circumference

}
