// DOCS: https://gin-gonic.com/docs/

package controllers

import (
	"go-server/db"
	"go-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := db.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	db.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	db.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	db.DB.First(&post, id)

	db.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	db.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
