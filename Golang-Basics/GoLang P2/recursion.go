package main 

// import "fmt"

func factorial(n int) int {
	if(n==0) {
		return 1
	}
	return n * factorial(n-1)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}