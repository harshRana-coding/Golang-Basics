package main

import "fmt"

type transformFunc func(int) int

func func1() {
	numbers := []int{1, 2, 3, 4, 5}
	dNumbers := doubleNumbers(&numbers,double)
	fmt.Println(dNumbers)
	tNumbers := doubleNumbers(&numbers,triple)
	fmt.Print(tNumbers)
}

func getTransformerFunction() transformFunc{
	return double
}

func doubleNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val)) 
	}
	return dNumbers
}

func double[T int | float64](number T) T {
	return number * 2
}
func triple[T int | float64](number T) T {
	return number * 3
}