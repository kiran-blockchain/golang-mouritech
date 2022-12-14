package main

import (
	"fmt"
	"time"
)

func myRoutine(ch chan int) {
	//in this example channel is sending the data to the method
	fmt.Println(100 + <-ch)
}
func main() {
	//channel will be created with a keyword chan
	c := make(chan int)
	// arrow is very important to identify the type of the channel
	// in this example channel c is of type int and will receive the data.
	go myRoutine(c)
	c <- 234
	time.Sleep(time.Second*10)
	fmt.Println("Value of channel ", c)
	fmt.Printf("Type of channel %T ", c)
}
