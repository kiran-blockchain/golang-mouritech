package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":9000"
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/contact", contact)
	log.Fatal(http.ListenAndServe(PORT, nil))

}

func index(resp http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(resp, "Welcome to web world of golang")
}
func home(resp http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(resp, "home page")
}
func contact(resp http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(resp, "contact us")
}
