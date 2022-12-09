package main

import (
	"fmt"
	"time"
)

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13} // arry of numbers
	nums2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 1000}
	c := make(chan int) //create an integer channel
	go sums(nums2, c)
	go sums(nums, c)
	x, y := <-c, <-c // Receive from c//receiver
	fmt.Println(x, y)
	

}
func sums(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
		fmt.Println(v)
	}
	//send sum to Channel//sender
	c <- sum
}
