package main

import "fmt"

func main() {
	var a = 7.98
	var p = &a
	var pp = &p

	fmt.Println("a = ", a)
	fmt.Println("address of a = ", &a)

	fmt.Println("p = ", p)
	fmt.Println("address of p = ", &p)

	fmt.Println("address of p in pp = ", pp)

	// Dereferencing a pointer to pointer
	fmt.Println("*pp = ", *pp)   //this will give the address of a
	fmt.Println("**pp = ", **pp) //this will give the data stored in a via p.
}
