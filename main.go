package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		if id == "1" {
			c.JSON(http.StatusOK, gin.H{"name": "PASHA"})
			return
		}

		if id == "2" {
			c.JSON(http.StatusOK, gin.H{"name": "TOXA"})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})

	r.Run()
}
