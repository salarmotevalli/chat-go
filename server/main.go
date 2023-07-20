package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
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
	if err := socket.Serve(); err != nil {
		log.Fatalf("socketio listen error: %s\n", err)
	}
}