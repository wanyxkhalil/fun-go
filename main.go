package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()
	os.Create("./fun_go.log")
	router := gin.Default()
	router.GET("/go/ifconfig", func(c *gin.Context) {
		c.String(200, c.ClientIP())
	})
	router.Run(":8001")
}
