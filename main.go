package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "errors"
)

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "up and running.")
}

func main() {
	router := gin.Default()
	router.GET("/health", health)
	router.Run("localhost:7777")
}