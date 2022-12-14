package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

type data struct {
	Body  []byte
	StatusCode int
	Error error
}
func main() {
	future := FetchInfoByPromise("https://restcountries.com/v2/all")
	// do many other things
	body := <-future
	fmt.Printf("response: %#v", string(body.Body))
	fmt.Printf("error: %#v", body.Error)
}
func FetchInfoByPromise(url string) <-chan data {
	c := make(chan data, 1)
	go func() {
		var body []byte
		var err error
		resp, err := http.Get(url)
		defer resp.Body.Close()
		
		body, err = ioutil.ReadAll(resp.Body)
		
		c <- data{Body: body, Error: err}
	}()
	return c
}
