package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shamysharma/task1/controler"
)

func UserRoute(router *gin.Engine) {
	router.GET("/task1/users", controler.GetEmp)
	router.POST("/task1/users", controler.PostEmp)
}
