package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sesamdev/kickstart-go-templ/internal/templates"
)

func HomePage(c *gin.Context) {
	rendered, err := templates.ExamplePage()
	if err != nil {
		c.String(500, "Error rendering page: %v", err)
		return
	}
	c.Data(200, "text/html; charset=utf-8", []byte(rendered))
}