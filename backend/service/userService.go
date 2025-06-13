package service

import (
	"backend/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		return
	}
	//这个地方应该是需要从数据库中进行查询
	userName, passWord := models.GetUserNameAndPwd(user.Username)
	if user.Username == userName && user.Password == passWord {
		log.Println("Login successful for user:", user.Username)
		c.JSON(200, gin.H{
			"data": "Login successful",
			"code": 200,
		})
	} else {
		fmt.Println("Invalid username or password")
		c.JSON(401, gin.H{"error": "Invalid username or password"})
	}
}

func UserRegister(c *gin.Context) {
	var userReg models.UserRegister
	if err := c.ShouldBindJSON(&userReg); err != nil {
		log.Println("Error binding JSON:", err)
		return
	}
	err := models.CreateUser(models.UserInformation{
		Username: userReg.Username,
		Password: userReg.Password,
		Email:    userReg.Email,
	})
	if err != nil {
		log.Println("Error creating user:", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, gin.H{
		"data": "User registered successfully",
		"code": 200,
	})
}

func UserInfo(c *gin.Context) {
	// models.GetUserByUsername()
}
