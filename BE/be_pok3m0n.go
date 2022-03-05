package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/labs/:labID/:status", func(c *gin.Context) {
		labID := c.Param("labID")
		status := c.Param("status")
		c.JSON(200, gin.H{
			"labID":     labID,
			"statusLab": status + " Lab",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
