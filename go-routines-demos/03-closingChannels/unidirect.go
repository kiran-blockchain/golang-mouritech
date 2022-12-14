package main

import "fmt"

func main(){
   sendCh:= make(chan  int)
   go senData(sendCh)
  fmt.Println(<-sendCh)
}

func senData(sendChannel chan<- int){
	sendChannel <- 10
}
