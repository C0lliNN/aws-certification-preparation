package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.Run(fmt.Sprintf(":%v", port))
}