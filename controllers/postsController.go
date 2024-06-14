package controllers

import (
	"go_crud/initializers"
	"go_crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get post body
	var body struct {
		Body  string `json:"body" binding:"required"`
		Title string `json:"title" binding:"required"`
	}

	// Bind JSON body to the struct and check for errors
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   "Request body parameters are missing or invalid",
		})
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	// Check for errors during the database operation
	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   result.Error.Error(),
		})
		return
	}

	// return the post
	c.JSON(200, gin.H{
		"success": true,
		"post":    post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	// Get all records
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"messge":  "Failed to get posts",
			"error":   result.Error.Error(),
		})
		return
	}

	// respond with posts
	c.JSON(200, gin.H{
		"success": true,
		"posts":   posts,
	})

}

func PostsShow(c *gin.Context) {
	var post models.Post
	// Get single record

	// get post by id from url
	id := c.Param("id")

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"messge":  "Failed to get post by id.",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"post":    post,
	})
}

func PostUpdate(c *gin.Context) {
	// get post id
	id := c.Param("id")

	// get post params
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// find and update post

	var post models.Post
	initializers.DB.First(&post, id)

	// Update attributes with `struct`, will only update non-zero fields
	result := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// return it

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"messge":  "Failed to get post by id.",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"post":    post,
	})
}

func PostDelete(c *gin.Context) {

	// get post by id from url
	id := c.Param("id")

	// DELETE FROM users WHERE id = 10;
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"messge":  "Failed to delete post.",
			"error":   result.Error.Error(),
		})
		return
	}

	// c.Status(200)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Successfuly deleted post",
	})

}
