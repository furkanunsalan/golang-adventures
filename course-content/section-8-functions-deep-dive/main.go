// variadic functions

package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)
	anaotherSum := sumup(1, numbers...) // sperator operator of js, directly

	fmt.Println(sum)
	fmt.Println(anaotherSum)

}

func sumup(startingValue int, numbers ...int) int { // like the seperator of js, this allows any number of integer variables, collectAll operator
	// sum := 0
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}

	return sum
}
