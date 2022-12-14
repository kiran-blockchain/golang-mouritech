package main

import "fmt"

// func main(){
// 	c:= make(chan int)
// 	c<-5
// }

func main(){
	messages:= make(chan string)
	go  func(){
		fmt.Println("I am Trying to Receive")
		<-messages
		fmt.Println("Receive Message")
	}()
	fmt.Println("Main trying to receive the data")
		<-messages
		fmt.Println("Receiving Mesages")
}