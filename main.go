package main

import (
	"fmt"
	"sort"
)

func main() {

	customers, err := readCustomersList("customers.txt")
	if err != nil {
		panic(err)
	}
	o := office{location{53.339428, -6.257664}}
	customersInRadius := o.getCustomersInRadius(customers, 100)
	printResult(customersInRadius)
}

func printResult(customers []*customer) {
	sort.Slice(customers, func(i, j int) bool {
		return customers[i].UserID < customers[j].UserID
	})
	fmt.Printf("Name\tUser ID\n")
	for i := 0; i < len(customers); i++ {
		fmt.Printf("%s\t%d\n", customers[i].Name, customers[i].UserID)
	}
}
