package controler

import (
	"github.com/gin-gonic/gin"
	"github.com/shamysharma/task1/config"
	"github.com/shamysharma/task1/models"
)

func GetEmp(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func PostEmp(c *gin.Context){
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}


