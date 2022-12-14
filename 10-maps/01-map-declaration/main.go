package main

import "fmt"

func main() {
	//var m map[string]int
	//fmt.Println(m)
	m := make(map[string] int)
	m["salary"]=2000 
	m["age"] =20
	m["employeeId"]=1

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println(m)
	}

	// The following statement will result in a runtime error
	// m["key"] = 100
}
