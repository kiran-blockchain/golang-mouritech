package main

import (
	"fmt"
	"net/http"

	//"runtime"

	"runtime"
	"time"
)

func routineCheck(from string) {
	start := time.Now() // Current Time
	for i := 0; i < 10000000000; i++ {
		//fmt.Println(from, ":", i)
	}
	elapsed := time.Since(start) // difference between start and end

	fmt.Println("Routine completed in " , elapsed)
}
func webCheck(_link string) {
	start := time.Now()
	_, err := http.Get(_link) // Hit the server and if the error is there say link down
	if err != nil {
		fmt.Println(_link, "might be down")
		return
	}
	elapsed := time.Since(start)
	fmt.Println(_link, "is up")
	fmt.Println("Time lapsed for execution ", elapsed, _link)
}

func main() {
	// regular function call
	//routineCheck("Direct Call")
	//runtime.GOMAXPROCS(16)
	start := time.Now()
	links := []string{
		"https://google.com",
		"https://tippintigers.com",
		"https://fb.com",
		"https://golang.org",
		"xy3..cood",
	}
	for _, link := range links {
		fmt.Println("Calling ", link)
		//This is how a normal function will be converted into a routine
		go webCheck(link)
	}

	//this is the way to call any function as a routine.
	for i := 0; i < 16; i++ {
		go routineCheck("Routine" + string(i))
	}
	elapsed := time.Since(start)
	time.Sleep(time.Second * 18)
	fmt.Println("Done")
	fmt.Println("Time lapsed for execution", elapsed)
}
