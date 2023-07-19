package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"chat/app/models"
	"chat/app/routes"
)

var (
	dbClient *mongo.Client
	ctx      context.Context
	cancel   context.CancelFunc
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
	defer cancel()

	wait = make(chan struct{}, 1)

	dbName := os.Getenv("DB_NAME")

	// Initialize app
	app := App{}
	models.Init(dbClient.Database(dbName), ctx)
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
