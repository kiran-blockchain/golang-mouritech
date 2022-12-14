package main

import "fmt"

func increment() func() int{
	i:=10
	return func () int{
		i++
		return i
	}
}

func main(){
	next:= increment()
	fmt.Println(next())
	fmt.Println(next())

}