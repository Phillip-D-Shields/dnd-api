package main

import (
	"example.com/gin-air-dnd/stats"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "kia ora, whanau!")
	})

	r.GET("/api/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	r.GET("/api/stats/generate", func(c *gin.Context) {
		characterStats := stats.GenerateStats()
		c.JSON(200, characterStats)
	})

	r.Run(":8080")
}
