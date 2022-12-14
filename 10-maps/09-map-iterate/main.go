package main

import "fmt"

func main() {
	var personAge = map[string]int{
		"Rajeev": 25,
		"James":  32,
		"Sarah":  29,
	}
	personAge["kiran"]= 35
	personAge["abhi"]= 40
	

	for name, age := range personAge {
		fmt.Println(name, age)
	}

}
