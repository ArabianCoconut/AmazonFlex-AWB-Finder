package main

import (
	"log"

	"github.com/ArabianCoconut/AmazonFlex-OrderTracker/src/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Println("Error loading .env file will load from system environment variables")

	} else {
		log.Println("Environment variables loaded successfully")
	}

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
	// loadenv()
	router := gin.Default()
	router.Static("./assets", "./assets")
	router.StaticFile("./favicon.ico", "./favicon.ico")
	router.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})

	// POST route to submit data is Working fine connected with the frontend
	// Database connection is not implemented
	router.POST("/api/submit", func(c *gin.Context) {
		var json struct {
			AWB      string `json:"awb" binding:"required"`
			DateTime string `json:"datetime"`
			Remark   string `json:"remark"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Data received successfully",
			"data":    json,
		})
		// Connect to the database and upload the data
		database.ConnectandUpload(json.AWB, json.DateTime, json.Remark)
	})

	return router
}
