package main

import (
	"fmt"
)

func myFunc(mychan chan string){
	for i:=0;i<4;i++{
		mychan <- "myChannel data"
	}
	close(mychan)
}


func main(){
	c:=make(chan string)
	go myFunc(c)
	// for {
	// 	res, ok:= <- c
	// 	if ok ==false{
	// 		fmt.Println("Channel Closed",ok)
	// 		break
	// 	}
	// 	fmt.Println("Channel is open ",res, ok)
	// }
	for res := range c{
		fmt.Println(res)
	}
}

