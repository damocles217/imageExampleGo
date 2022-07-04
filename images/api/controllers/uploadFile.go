package controllers

import (
	"context"
	"net/http"
	"path/filepath"
	"time"

	"github.com/damocles217/images_service/images/api/utils"
	"github.com/damocles217/images_service/images/core/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func UploadFile(imageCollection *mongo.Collection, cacheUser *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		_, id, cUser := utils.Auth(c)

		cacheUser.Del(context.TODO(), "images_"+id)

		file, err := c.FormFile("image")

		// Upload the file to specific dst.
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is received",
			})
			return
		}

		// Retrieve file information
		extension := filepath.Ext(file.Filename)
		// Generate random file name for the new uploaded file so it doesn't override the old file with same name
		newFileName := uuid.New().String() + extension

		pathName := "./uploads/" + cUser + "/"

		utils.CreateDir(pathName)

		image := models.Image{
			UrlPhoto:  cUser + "/" + newFileName,
			CreatedAt: time.Now(),
			Owner:     id,
		}

		image.UpdatedAt = image.CreatedAt

		imageCollection.InsertOne(context.TODO(), image)

		// The file is received, so let's save it
		if err := c.SaveUploadedFile(file, pathName+newFileName); err != nil {

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}

		// File saved successfully. Return proper result
		c.JSON(http.StatusOK, gin.H{
			"message": "Your file has been successfully uploaded.",
		})
	}
}
