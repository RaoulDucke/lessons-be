package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/RaolDucke/lessons-be/db"
	"github.com/RaolDucke/lessons-be/handler"
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

	h := handler.New(db.New())

	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, h.GetProducts())
	})
	r.POST("/products", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Internal error")
			return
		}
		product := new(handler.Product)
		err = json.Unmarshal(jsonData, product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Internal error")
			return
		}
		err = h.AddProduct(product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Internal error")
			return
		}
	})

	r.Run()
}
