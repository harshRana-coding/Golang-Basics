package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func Init() {
	var productNames [4]string
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	productNames[2] = "Race Gurram"
	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[3])
}
//dynamic
func Init2() {
	prices := []float64{33.99,44.00}
	fmt.Println(len(prices),cap(prices));
	prices = append(prices, 5.99,6.77)
	fmt.Println(prices);
	fmt.Println(len(prices),cap(prices));
}

func practice() {
	//1
	hobbies := [3]string{"Coding","Video Games","Volleyball"}
	fmt.Println(hobbies)

	//2
	fmt.Printf("My first hobby is %v\n", hobbies[0])
	fmt.Printf("My new hobbies are %v\n", hobbies[1:])

	//3
	oldHobbies := hobbies[:2]
	fmt.Println(oldHobbies)

	//4
	oldHobbies = oldHobbies[1:]
	fmt.Println(oldHobbies)

	//5
	goals := []string{"Learn Go", "Build a API"}
	fmt.Println(goals)

	//6
	goals[1] = "Build a REST API"
	goals = append(goals, "Complete the WABA Task")
	fmt.Println(goals)

	//7
	product1 := Product{
		title: "Hair Wash",
		id: "123",
		price: 22.4,
	}
	product2 := Product{
		title: "Hair Cream",
		id: "124",
		price: 13.4,
	}
	products := []Product{product1,product2}
	product3 := Product{
		title: "Hair Gel",
		id: "125",
		price: 10.99,
	}
	products = append(products, product3)
	fmt.Println(products)
}
