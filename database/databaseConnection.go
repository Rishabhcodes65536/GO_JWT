package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func close(client *mongo.Client, ctx context.Context,
           cancel context.CancelFunc){
            
    // CancelFunc to cancel to context
    defer cancel()
     
    // client provides a method to close 
    // a mongoDB connection.
    // defer func(){
     
    //     // client.Disconnect method also has deadline.
    //     // returns error if any,
    //     if err := client.Disconnect(ctx); err != nil{
    //         panic(err)
    //     }
    // }()
}

func connect(uri string)(*mongo.Client, context.Context, 
                          context.CancelFunc, error) {
                           
    // ctx will be used to set deadline for process, here 
    // deadline will of 30 seconds.
    ctx, cancel := context.WithTimeout(context.Background(), 
                                       30 * time.Second)
     
    // mongo.Connect return mongo.Client method
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error{
 
    // mongo.Client has Ping to ping mongoDB, deadline of 
    // the Ping method will be determined by cxt
    // Ping method return error if any occurred, then
    // the error can be handled.
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected successfully")
    return nil
}

func DBInstance()  *mongo.Client{
	err := godotenv.Load(".env")
	if err!=nil {
		log.Fatal("Error loading .env file")
	}
	Mongodb := os.Getenv("MONGODB_URL")

	client,ctx,cancel,err :=connect(Mongodb)


	if err!=nil{
		log.Fatal("Error")
	}

	 // Release resource when the main
    // function is returned.
    defer close(client, ctx, cancel)
     
    // Ping mongoDB with Ping method
    ping(client, ctx)

	fmt.Println("Connected to MongoDB!")

	return client
}

var Client *mongo.Client =DBInstance()

func OpenCollection(client *mongo.Client, collectionName string)	*mongo.Collection{
	var collection *mongo.Collection=client.Database("cluster0").Collection(collectionName)
	return collection
}