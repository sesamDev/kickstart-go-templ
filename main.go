package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sesamdev/kickstart-go-templ/internal/handlers"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./dist")

	r.GET("/", handlers.HomePage)
	r.Run(":8080")
}
