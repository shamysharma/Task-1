package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shamysharma/task1/config"
	"github.com/shamysharma/task1/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")

}
