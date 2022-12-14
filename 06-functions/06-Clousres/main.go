package main

import "fmt"

func main(){
   x, y :=adder(),adder()

   for i:=0; i<5; i++ {
	  fmt.Println(x(i))
	  fmt.Println(y(-2*i))
   }
   fmt.Println(x(32))
   
}

func adder() func(int) int{
	 sum:=0
	 return func(x int) int{
		sum +=x
		return sum
	 }
}