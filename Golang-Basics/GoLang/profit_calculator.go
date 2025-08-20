package main

import (
	"errors"
	"fmt"
)

func profitCalculator(){
	// var revenue, expenses, taxRate float64
	revenue, err1 := takeInput("Enter the revenue : ")
	// if err!=nil{
	// 	fmt.Println(err);
	// 	return;
	// }
	expenses, err2 := takeInput("Enter the expenses : ")
	// if err!=nil{
	// 	fmt.Println(err);
	// 	return;
	// }
	taxRate, err3 := takeInput("Enter the taxRate : ")
	if err1!=nil || err2!=nil || err3!=nil{
		fmt.Println(err1);
		return;
	}

	var ebt, profit, ratio = calc(revenue, expenses, taxRate);
	printFloat(ebt);
	printFloat(profit);
	printFloat(ratio);
}

func printFloat(val float64){
	fmt.Printf("%.1f\n", val)
}

func takeInput(text string) (float64, error){
	var val float64
	fmt.Print(text)
	fmt.Scan(&val)

	if (val <= 0){
		return 0, errors.New("Value must be a positive number.")
	}

	return val, nil
}

func calc(revenue, expenses, taxRate float64)(ebt, profit, ratio float64){
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt/profit
	return
}