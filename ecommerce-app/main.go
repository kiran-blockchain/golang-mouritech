package main

import (
	"ecommerce-app/route"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Main Application Starts")
	//Loading Environmental Variable
	loadenv()

	log.Fatal(route.RunAPI(":8090"))

}
