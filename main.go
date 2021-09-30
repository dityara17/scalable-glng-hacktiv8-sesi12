package main

import (
	"go-jwt-glng-aditya/database"
	"go-jwt-glng-aditya/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
