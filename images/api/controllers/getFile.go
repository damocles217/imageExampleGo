package controllers

import "github.com/gin-gonic/gin"

func GetFile(c *gin.Context) {
	user := c.Param("user")
	id := c.Param("id")
	c.Writer.Header().Set("Content-Type", "image/*")

	c.File("./uploads/" + user + "/" + id)
}
