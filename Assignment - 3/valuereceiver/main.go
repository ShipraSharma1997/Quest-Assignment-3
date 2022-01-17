package main

import "fmt"

func main() {
	createCustomer()
}

func createCustomer() {

	// Initialise using make function
	customerMap := make(map[int]int)
	customerMap[1] = 100

	for _, value := range customerMap {
		fmt.Println("Customer Id is: ", value)
	}

	//Initialise using map literal
	customers := map[int]int{
		1: 200,
		2: 400,
	}

	//for delete
	delete(customers, 2)

	//get one element
	fmt.Println("Customer Id is: ", customers[1])

	//get all elements
	fmt.Println("Customer Id is: ", customers)

	for _, value := range customers {
		fmt.Println("Customer Id is: ", value)
	}

}
