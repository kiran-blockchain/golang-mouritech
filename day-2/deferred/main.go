package main

import "fmt"

func main() {
	fmt.Println("1")
	defer myDemo()
	 myDemo2()
	fmt.Println("3")
  }
   
  func myDemo() {
	fmt.Println("2")
  }
   
  func myDemo2() {
	defer myDemo3()
  }
   
  func myDemo3() {
	fmt.Println("2.1")
  }
  