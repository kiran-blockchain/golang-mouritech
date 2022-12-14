package main

import (
    "fmt"
    "time"
)
func worker(id int, jobs <-chan int, results chan<-int) {
    for j := range jobs {
       // fmt.Println("worker", id, "started  job", j)
        //time.Sleep(time.Second)
        //fmt.Println("worker", id, "finished job", j)
        results <- j * 2//
    }
}
func main() {
    start:=time.Now();
    const numJobs = 10000
    jobs := make(chan int, numJobs)// belt 1
    results := make(chan int, numJobs)// belt 2
    for w := 1; w <= 300; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j// sending the data on to the jobs channel
    }
    close(jobs)

    for a := 1; a <= numJobs; a++ {
        <-results// i am receiving the data on the results channel
   }
   elapsed:=time.Since(start)
   fmt.Printf("Time for execution %s ",elapsed)
}