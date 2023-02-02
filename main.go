package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/RaolDucke/lessons-be/db"
	"github.com/RaolDucke/lessons-be/handler"
	"github.com/RaolDucke/lessons-be/handlerUsers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "usr"
	password = "pwd"
	dbname   = "products"
)

func main() {
	ctx := context.Background()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	repository := db.New(database)

	h := handler.New(repository)
	u := handlerUsers.New(repository)

	r.GET("/products", func(c *gin.Context) { h.GetProducts(ctx, c) })
	r.POST("/products", func(c *gin.Context) { h.AddProduct(ctx, c) })
	r.PUT("/products", func(c *gin.Context) { h.UpdateProduct(ctx, c) })
	r.POST("/users", u.AddUser)
	r.GET("/users", u.GetUsers)

	r.Run()
}
