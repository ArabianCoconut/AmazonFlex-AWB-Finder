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
	router := runGin()
	router.Run(":8080")

}

// runGin initializes and returns a Gin engine with predefined routes and handlers.
// 
// Routes:
// - GET /: Serves the index.html file.
// - POST /api/submit: Accepts JSON data with AWB, DateTime, and Remark fields, validates it, and responds with a success message. The data is then uploaded to the database.
// - GET /api/fetch: Fetches data from the database and returns it as JSON.
// - GET /portal: Serves the portal.html file.
// - POST /api/delete: Accepts JSON data with an AWB field, validates it, and responds with a success message. The corresponding data is then deleted from the database.
//
// Static Files:
// - ./assets: Serves static files from the ./assets directory.
// - ./favicon.ico: Serves the favicon.ico file.
//
// Returns:
// - *gin.Engine: The initialized Gin engine.
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
			"message": "Data saved successfully",
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
