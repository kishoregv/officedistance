package main

import (
	"bufio"
	"os"
)

func readCustomersList(filename string) ([]*customer, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)

	customers := []*customer{}
	for scanner.Scan() {
		bytes := scanner.Bytes()
		var c customer

		if err := c.unmarshal(bytes); err != nil {
			return nil, err
		}
		customers = append(customers, &c)
	}
	return customers, nil
}
