package main

import (
	"events-booking/db"
	"events-booking/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Events Booking API")
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
