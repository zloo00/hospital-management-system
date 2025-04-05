package main

import (
	"github.com/aluazholdykan/hospital-management-system/config"
	"github.com/aluazholdykan/hospital-management-system/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ensure ConnectDatabase is called before setting up routes
	config.ConnectDatabase()

	// Register routes after database connection is established
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
