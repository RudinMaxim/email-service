package router

import (
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/RudinMaxim/email-service.git/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/RudinMaxim/email-service.git/docs"
)

var r *gin.Engine

func InitRouter() error {
	r = gin.Default()

	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatal("Error loading .env file")
	}

	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return err
	}
	r.Use(CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/send-email", service.SendEmailHandler)

	return nil
}

func Start(addr string) error {
	return r.Run(addr)
}
