package main

import "fmt"

func myfunc(channel chan string) {
	for i := 0; i < 10; i++ {
		channel <- "Go Lang is Awesome "
	}
	fmt.Println("Lengh of the channel is", len(channel))
	close(channel)
}
func main() {
	c := make(chan string, 5)
	go myfunc(c)
	//my func is acting as the data publisher
	fmt.Println("Capacity of the channel is", cap(c))
	counter := 0
	for {
		res, status := <-c
		counter++
		if !status {
			fmt.Println("Channel is closed", status)
			break
		}
		fmt.Println("Channel is open and data is", res, status)
	}
}
