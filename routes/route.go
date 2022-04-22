package routes

import (
	"bytebank-api/controller"
	"bytebank-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
		AllowCredentials: true,
	}))

	main := r.Group("/api/v1")
	{
		user := main.Group("/user")
		{
			user.POST("/register", controller.Register)
			user.POST("/login", controller.Login)
		}

		transferencias := main.Group("/transferencias", middlewares.Auth())
		{
			transferencias.GET("", controller.GetTransferencias)
			transferencias.POST("", controller.InsertTransferencia)
		}
	}
	r.Run()
}
