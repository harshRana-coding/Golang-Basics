package main

import (
	// "fmt"
	// "pricecalculator/v1/filemanager"
	"fmt"
	"pricecalculator/v1/cmdmanager"
	"pricecalculator/v1/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100))
		cm := cmdmanagergo.New();
		priceJob := prices.NewTaxIncludedPriceJob(cm,taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Error processing Job:", err)
		}
	}

}
