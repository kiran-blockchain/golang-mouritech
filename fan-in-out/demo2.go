package main

import (
	"fmt"
	
	"time"
)


func main(){
	ch :=make(chan int,1)
    
	go receiver(ch,1)
	
	time.Sleep(5*time.Second)
	go receiver(ch,2)
	go sender(ch,100)
	fmt.Println("I am done")
}

func receiver(ch chan int,num int){
	
	x:= <-ch

	fmt.Println("Receiving Data",num)
	fmt.Println(x)
 }
 func sender(ch chan int,data int){
	ch<-data
 }

