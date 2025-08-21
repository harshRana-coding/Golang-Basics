package prices

import (
	"fmt"

	"pricecalculator/v1/conversion"
	"pricecalculator/v1/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"`
	TaxRate           float64            `json:"taxRate"`
	InputPrices       []float64          `json:"inputPrices"`
	TaxIncludedPrices map[string]string `json:"taxIncludedPrices"`
}

func (job *TaxIncludedPriceJob) loadPrices() error {
	lines,err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	prices,err := conversion.StringToFloat(lines)
	if err != nil {
		return err
	}
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error{
	err := job.loadPrices()
	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1+job.TaxRate)
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	return job.IOManager.WriteLines(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager,taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:  iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
