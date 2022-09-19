package main

import (
	"fmt"
	"github.com/deltamc/otus-social-networks-users/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	if len(os.Getenv("RUN_IN_DOCKER")) > 0 {
		return
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	routes.Public()
	routes.Auth()
	fmt.Println(os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
