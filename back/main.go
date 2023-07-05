package main

import (
    "log"
    "os"
    "fmt"
	"net/http"

    "github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal(err.Error())
	}
  
	appPort := os.Getenv("APP_PORT")
	routs := set_up_routes()
	

	http.ListenAndServe(fmt.Sprintf(":%s", appPort), routes)
}