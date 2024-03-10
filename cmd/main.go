package main

import (
	"github.com/RudinMaxim/email-service.git/router"
)

// @title API сервиса Email
// @version 1.0
// @description Это API сервер для сервиса email.

// @host localhost:8080
// @BasePath /

func main() {
	err := router.InitRouter()
	if err != nil {
		return
	}

	router.Start("0.0.0.0:8080")
}
