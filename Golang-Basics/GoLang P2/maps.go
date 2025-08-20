package main 

import "fmt"

type floatMap map[string]float64

func map1() {
	websites := map[string]string{
		"Google" : "https://google.com",
		"Amazon Web Services" : "https://aws.com",
	}
	fmt.Println(websites["Google"])
	websites["LinkedIn"] = "https://linkedin.com"
	delete(websites, "Google")
	fmt.Println(websites)
}

func slice() {
	userNames := make([]string, 2, 5)
	userNames[0] = "Kumar"
	userNames = append(userNames,"max")
	userNames = append(userNames,"manuel")
	fmt.Println(userNames);
	fmt.Println(len(userNames), cap(userNames));

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["React"] = 4.6
	courseRatings["Angular"] = 4.4
	courseRatings["Angular"] = 4.4
	courseRatings["English"] = 4.4
	fmt.Println(courseRatings)

	for index, value := range userNames{
		fmt.Println(index, value)
	}
	for key, value := range courseRatings{
		fmt.Println(key, value)
	}
}