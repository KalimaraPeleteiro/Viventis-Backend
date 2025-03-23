package main

import (
	"viventis/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDatabase()

	r := gin.Default()
	r.Run(":8080")
}
