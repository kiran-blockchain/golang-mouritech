package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"context"	
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var DB *mongo.Client = ConnectDB()

func GetEnvData(key string) string{
		err := godotenv.Load() //loads .env file in to the project environment
		if err != nil {
			fmt.Println(err)
			fmt.Println("unable to load envoriment")
			return (":9000") //return default :3000 as port number
		}
		return os.Getenv(key)
}

func ConnectDB() *mongo.Client {
		//Get the instance of the datbase client.
		client, err := mongo.NewClient(options.Client().ApplyURI(GetEnvData("MONGODBURI")))
		if err != nil {
			fmt.Println("Error in Connecting to Mongo db Database")
		}
		fmt.Println("Database getting connected")
		//Set the context to run the client in the background
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
		//Connect the client and run the background
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		//ping the database
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to MongoDB")
	
		return client
}
	
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
		collection := client.Database("productsdb").Collection(collectionName)
		return collection
}



