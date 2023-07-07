package main

import (
	"chat/app/routes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbClient *mongo.Client
	ctx      context.Context
)

type App struct {
	conn *mongo.Client
}

func init() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the mongo
	dbClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	// Check mongo ping
	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}
}

// waiter for serve goroutine
var wait chan struct{}

func main() {

	defer dbClient.Disconnect(ctx)

	wait = make(chan struct{}, 1)

	// Initialize app
	app := App{
		conn: dbClient,
	}

	// Serve app
	go app.serve()

	// Wait until end serving
	<-wait
}

func (app *App) serve() {
	// Get routes
	appRoutes := routes.Setup()
	appPort := os.Getenv("APP_PORT")

	// Start serving app
	err := http.ListenAndServe(fmt.Sprintf(":%s", appPort), appRoutes)
	if err != nil {
		panic(err.Error())
	}

	wait <- struct{}{}
}
