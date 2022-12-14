package main

import (
	"fmt"
	"time"
)

func main() {
	x("direct")
	go x("routine")
	//inline go routine
	go func(msg string){
		fmt.Println(msg)
	}("goingg")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func x(y string) {
	for i := 0; i < 3; i++ {
		fmt.Println(y ,":",i)
	}
}
