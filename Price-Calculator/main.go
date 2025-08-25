package main

import (
	// "fmt"
	"pricecalculator/v1/filemanager"
	"fmt"
	// "pricecalculator/v1/cmdmanager"
	"pricecalculator/v1/prices"
)

func main() {
	taxRates := []float64{0, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100))
		// cm := cmdmanagergo.New();
		priceJob := prices.NewTaxIncludedPriceJob(fm,taxRate)
		go priceJob.Process(doneChans[index],errorChans[index])
		// if err != nil {
		// 	fmt.Println("Error processing Job:", err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <- errorChans[index]:
			if err != nil {
				fmt.Println("Error processing Job:", err)
			}
		case <- doneChans[index]:
			fmt.Printf("Job %d completed successfully\n", index)
		}
	}

}
