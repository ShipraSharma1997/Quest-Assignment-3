package customer

import "fmt"

//Initialise struct
type Customer struct {
	CustomerID int
}

//CreateCustomer
func createCustomer() {
	customer1 := Customer{
		CustomerID: 100,
	}

	customer2 := Customer{
		CustomerID: 200,
	}

	//Initialise Map by make keyword
	customersMap := make(map[int]Customer)
	customersMap[1] = customer1
	customersMap[2] = customer2

	fmt.Printf(fmt.Sprintf("\n Map of struct: %+v", customersMap))
}
