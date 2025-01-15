package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, `
     ____    _______
	|____|  |**___**|       
	______  |*|___|*|
   |******| |*******|
   |******| |*******|
   |******| |**/ \**|
   |******| |__| |__|
   |______| 
`)
	})
	r.Run(":8080")
}
