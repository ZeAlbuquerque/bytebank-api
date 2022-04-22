package controller

import (
	"bytebank-api/database"
	"bytebank-api/models"
	"bytebank-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const SecretKey = "secret"

func Register(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	password := services.SHA256Encoder(data["password"])

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var login models.Login
	var user models.User

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Where("email = ?", login.Email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}

func User(c *gin.Context) {

}
