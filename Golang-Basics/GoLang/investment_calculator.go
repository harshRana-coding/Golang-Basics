package main

import (
	"fmt"
)

func main() {
	// const inflationRate = 6.5
	// var investmentAmount, expectedReturnRate float64
	// var years float64

	// //fmt.Print("Enter the investment Amount : ")
	// outputText("Enter the investment Amount : ")
	// fmt.Scan(&investmentAmount)

	// // fmt.Print("Enter the return rate : ")
	// outputText("Enter the return rate : ")
	// fmt.Scan(&expectedReturnRate)

	// // fmt.Print("Enter the Investment Period: ")
	// outputText("Enter the Investment Period: ")
	// fmt.Scan(&years)

	// var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// formattedFV := fmt.Sprintf("Future Value : %0.1f\n", futureValue)
	// formattedFRV := fmt.Sprintf("Future Real Value : %0.1f\n", futureRealValue)
	// // fmt.Println("Future Value : ", futureValue)
	// // fmt.Printf("Future Value %.0f\nFuture Real Value: %.1f", futureValue, futureRealValue)
	// // fmt.Println("Future Real Value : ", futureRealValue)
	// fmt.Print(formattedFV,formattedFRV)
	// profitCalculator();
	// bank()
	// student.Init()
	// note.Init()
	generic()
}

func outputText(text string){
	fmt.Print(text)
}
