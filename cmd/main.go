package main

import (
	"github.com/RudinMaxim/email-service.git/router"
)

func main() {

	err := router.InitRouter()
	if err != nil {
		return
	}

	router.Start("0.0.0.0:8080")
}
