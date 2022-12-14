package main

import (
	"fmt"
	//"time"
)

func main() {
	
	ch := make(chan int)
	
    x:=<-ch
	ch <- 1
	fmt.Println(x)

	//go write(ch)
	//time.Sleep(2*time.Second)
	// for v:= range ch{
	// 	fmt.Println("Reading Value from channel-",v)
	// 	//time.Sleep(2*time.Second)
	// }
}

// func write(ch chan int){
// 	// for i:=0;i<100;i++{
// 	// 	ch <- i
// 	// 	fmt.Println("Successfully written to channel:",i)
// 	// }
// 	// close(ch)
// }
