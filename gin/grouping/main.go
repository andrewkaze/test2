package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func readEndpoint(context *gin.Context) {
	context.JSON(200, gin.H{"endpoint": "read"})
}

func submitEndpoint(context *gin.Context) {
	context.String(http.StatusOK, "Submitted")
}

func loginEndpoint(context *gin.Context) {
	context.String(http.StatusOK, "Logged In")
}
