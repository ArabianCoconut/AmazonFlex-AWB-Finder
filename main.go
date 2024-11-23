package main

import (

	"github.com/gin-gonic/gin"

)

func main() {
	gin.SetMode(gin.DebugMode)
	r := runGin()
	r.Run(":8080")
}

// runGin initializes and returns a Gin engine with predefined routes and handlers.
// It serves static files from the "./assets" directory and a favicon from "./favicon.ico".
// The root route ("/") serves the "index.html" file.
// The "/api/submit" POST route accepts a JSON payload with "awb", "date", and "time" fields,
// validates the payload, and responds with a success message and the received data.

func runGin() *gin.Engine {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})

	// POST route to submit data is Working fine connected with the frontend
	// Database connection is not implemented
	router.POST("/api/submit", func(c *gin.Context) {
		var json struct {
			AWB  string `json:"awb" binding:"required"`
			DateTime string `json:"datetime"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"data":    json,
		})
	})

	return router
}
