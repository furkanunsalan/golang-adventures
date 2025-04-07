package main

import "fmt"

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

// Not meaningful because you would have known how many days you will have in advance
// type TemperatureData struct {
// 	day1 float64
// 	day2 float64
// }

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:] // 1: or :3 is possible as in most languages but negative index is not allowed like [-1]
	featuredPrices[0] = 199.99 // editing this also edits the element in the original array because slices are actually windows to the orginial array
	// so when you create a slice, the array is not copied thus saves from the memory
	highlightedPrices := featuredPrices[:1]

	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 1, 3 (because cap extends towrds the end)
	// len: number of elements inside an array
	// cap: number of elements in the array which the slice is based on, can extend towards the end but not front
	// for featuredPrices := prices[1:] -> cap() is 3 because there are 3 elements in the ancestor array towards the end, not 4 because you can't take the element in the front
	// for highlightedPrices := featuredPrices[:1] -> cap is again 3 because there are still elements that can be chosen towards the end

	highlightedPrices = highlightedPrices[:3]
	// this can work because you can choose more towards the right of the array
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 3, 3
}
