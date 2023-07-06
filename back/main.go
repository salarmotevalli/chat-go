package main

import (
	"chat/app/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	appRoutes := routes.Setup()
	appPort := os.Getenv("APP_PORT")

	err = http.ListenAndServe(fmt.Sprintf(":%s", appPort), appRoutes)
	if err != nil {
		panic(err.Error())
	}
}
