package main

import "github.com/gin-gonic/gin"


// main is the entry point of the application. It sets the Gin framework to release mode,
// creates a default Gin router, serves a static file "index.html" at the root URL path,
// and starts the HTTP server on port 8080.
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFile("/index.html", "./index.html")
	r.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})
	r.Run(":8080")
}