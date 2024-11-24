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

	gin.SetMode(gin.ReleaseMode)
	r := runGin()
	r.Run(":8080")

}

// runGin initializes and returns a Gin engine with predefined routes and handlers.
// It serves static files from the "./assets" directory and a favicon from "./favicon.ico".
// The root route ("/") serves the "index.html" file.
// The "/api/submit" POST route accepts JSON data with "awb", "datetime", and "remark" fields,
// validates the input, and responds with a success message and the received data.
// If the input is valid, it connects to the database and uploads the data.
// The "/api/submit" POST route accepts a JSON payload with "awb", "date", and "time" fields,
// validates the payload, and responds with a success message and the received data.

func runGin() *gin.Engine {
	router := gin.Default()
	router.Static("./assets", "./assets")
	router.StaticFile("./favicon.ico", "./favicon.ico")
	router.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})

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
		database.ConnectAndUpload(json.AWB, json.DateTime, json.Remark)
	})

	router.GET("/api/fetch", func(c *gin.Context) {
		data := database.ConnectAndFetch()
		if data == nil {
			c.JSON(400, gin.H{"error": "Error fetching data"})
			return
		}
		c.JSON(200, gin.H{
			"data": data,
		})
	})

	router.GET("/portal", func(c *gin.Context) {
		c.File("./portal.html")
	})

	router.POST("/api/delete", func(c *gin.Context) {
		var json struct {
			AWB string `json:"awb" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": json.AWB + " deleted successfully"})
		database.ConnectAndDelete(json.AWB)
		
	})

	return router
}
