package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) // type, initial size, capacity (intended length)
	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	// courseRatings := make(map[string]float64, 3) // type, capacity (intended length)
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7
	courseRatings["node"] = 4.7

	courseRatings.output()
	// fmt.Println(courseRatings)

	for index, value :=  range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings{
		fmt.Println(key, value)
	}
}