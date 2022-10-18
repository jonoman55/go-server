package main

import (
	"fmt"
	"go-server/cfg"
	"go-server/controllers"
	"go-server/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	cfg.LoadEnvVars()
	db.ConnectDB()
}

func main() {

	router := gin.Default()

	router.SetTrustedProxies([]string{"*"})

	router.Use(cors.Default())

	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPost)
	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	fmt.Println("Starting Server...")

	router.Run()

}
