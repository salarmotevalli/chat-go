package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"chat/app/models"
	"chat/app/routes"
)

var (
	mongoClient *mongo.Client
	redisClient *redis.Client
	ctx         context.Context
	cancel      context.CancelFunc
)

type App struct{}

func init() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize ctx
	ctx, cancel = context.WithCancel(context.Background())

	// Connect to the mongo
	mongoClient, err = connectToMongo()
	if err != nil {
		log.Fatal("Could not connect to Mongo: ", err)
	}

	redisClient, err = connectToRedis()
	if err != nil {
		log.Fatal("Could not connect to Redis: ", err)
	}

}

// waiter for serve goroutine
var wait chan struct{}

func main() {
	defer mongoClient.Disconnect(ctx)
	defer redisClient.Close()
	defer cancel()

	wait = make(chan struct{}, 1)

	dbName := os.Getenv("DB_NAME")

	// Initialize app
	app := App{}
	models.Init(mongoClient.Database(dbName), ctx)

	engine := gin.New()
	socket := socketio.NewServer(nil)

	routes.Setup(engine, socket)
	appPort := os.Getenv("APP_PORT")

	// Serve app
	go app.serve(engine, socket, appPort)

	// Wait until end serving
	<-wait
}

func (app *App) serve(engine *gin.Engine, socket *socketio.Server, appPort string) {
	// go runSocketServer(socket)
	go runSocketServer(socket)

	defer socket.Close()

	engine.GET("/socket.io/*any", gin.WrapH(socket))
	engine.POST("/socket.io/*any", gin.WrapH(socket))
	engine.StaticFS("/public", http.Dir("./asset"))

	defer socket.Close()

	// Start serving app
	err := http.ListenAndServe(fmt.Sprintf(":%s", appPort), engine)
	if err != nil {
		panic(err.Error())
	}

	wait <- struct{}{}
}

func runSocketServer(socket *socketio.Server) {
	// Serve socket server
	if err := socket.Serve(); err != nil {
		log.Fatalf("socketio listen error: %s\n", err)
	}
}

/*
	These functions are for connecting to database
	they will try to connect 5 times and return
	connection and error
*/


func connectToMongo() (*mongo.Client, error) {
	var err error
	var client *mongo.Client

	for _ = range [5]struct{}{} {

		log.Println("Connecting to Mongo...")

		client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

		// Check mongo ping
		err = client.Ping(ctx, readpref.Primary())

		if err == nil {
			log.Println("Connected to Mongo.")
			return client, err
		}

		log.Println("backing off...")
		time.Sleep(time.Second)
	}

	return nil, err

}

func connectToRedis() (*redis.Client, error) {
	var err error
	var client *redis.Client

	for _ = range [5]struct{}{} {

		log.Println("Connecting to Redis...")

		client = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "salar",
			DB: 0,
		})

		// Check redis ping
		pong, err := client.Ping(ctx).Result()
		if err == nil && pong == "PONG" {
			log.Println("Connected to Redis.")
			return client, nil
		}

		log.Println("backing off...")
		time.Sleep(time.Second)
	}

	return nil, err

}
