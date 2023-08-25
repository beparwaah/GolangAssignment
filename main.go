package main

import (
	"github.com/beparwaah/assignmentVegapay/config"
	"github.com/beparwaah/assignmentVegapay/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.CreditRoutes(router)
	router.Run(":8080")
}
