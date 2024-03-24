package main

import (
	"fmt"
	"prices-calculator/filemanager"
	"prices-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	// prices :=

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		err := priceJob.Process()
		if err != nil {
			fmt.Println("Couldnt process job")
			fmt.Println(err)
		}

	}

}
