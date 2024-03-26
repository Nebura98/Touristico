package main

import (
	"fmt"
	"log"

	dbconfig "main/internal/adapters/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("development.env")

	if err != nil {
		fmt.Println(err)

		log.Fatal("Error loading .env file")
	}

	dbconfig.InitDatabase()

	fc, _ := dbconfig.Pool.GetConnection()
	println("First connection: %s", fc)

	sc, _ := dbconfig.Pool.GetConnection()
	println("Second connection: %s", sc)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Run()
}
