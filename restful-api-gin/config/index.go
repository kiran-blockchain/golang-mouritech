package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


func GetEnvData(key string) string{
		err := godotenv.Load() //loads .env file in to the project environment
		if err != nil {
			fmt.Println(err)
			fmt.Println("unable to load envoriment")
			return (":9000") //return default :3000 as port number
		}
		return os.Getenv(key)
	}
