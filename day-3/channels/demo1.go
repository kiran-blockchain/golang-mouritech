package main

import (
	"fmt"
	//"time"
)


func main(){
	messages := make(chan string)//Bi-directional channel
 
	
   go func(){
	    //time.Sleep(5*time.Second)
		
		messages <-"Ping" //sending the data to the channel
		fmt.Println("I am done")
	}()
  
	// when the data is available on the channel will be subscribed here
    //msg:= <-messages 
	for i:=0;i<1000000;i++{}
	fmt.Println("check")
}