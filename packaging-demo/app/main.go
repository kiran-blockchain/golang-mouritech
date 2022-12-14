package main

import (
	"fmt"
	"hcl-golang/packaging-demo/numbers"
)

func main(){
	check:= numbers.IsPrime(123)
	fmt.Println(check)
}