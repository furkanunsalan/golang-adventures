package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:]) // or hobbies[1:3]

	// 3
	mainHobbies := hobbies[:2] // or hobbies[0:2]
	fmt.Println(mainHobbies)

	// 4
	fmt.Println(cap(mainHobbies))  //extends to right
	mainHobbies = mainHobbies[1:3] // expends need the last index specified explicityly because mainHobbies[1:] won't work
	fmt.Println(mainHobbies)

	// 5
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// 6
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)

	// 7
	products := []Product{
		// Product{
		// 	"first-product", 
		// 	"A first product", 
		// 	12.99,
		// },
		// Product{
		// 	"second-product", 
		// 	"A second product", 
		// 	129.99,
		// }}

		// same as the following because the slice only can take "Product" typed entries
		{
			"first-product", 
			"A first product", 
			12.99,
		},
		{
			"second-product", 
			"A second product", 
			129.99,
		}}

	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A thrid product",
		15.99,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}
