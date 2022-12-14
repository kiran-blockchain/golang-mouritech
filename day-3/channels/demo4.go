package main

import (
	"fmt"
	"time"
	"sync"
)

var x = 0

func main() {

	go increment()
	go increment2()

	time.Sleep(2 * time.Second)
	fmt.Println(x)

}

func increment() {
	fmt.Println("increment", x)
	var mu sync.Mutex
	mu.Lock()
	x= x+1
	mu.Unlock()

}
func increment2() {
	fmt.Println("increment2", x)
	var mu sync.Mutex
	mu.Lock()
	x= x+2
	mu.Unlock()
}
