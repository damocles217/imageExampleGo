package controllers

import (
	"context"
	"net/http"
	"path/filepath"
	"time"

	"github.com/damocles217/images_service/images/api/utils"
	"github.com/damocles217/images_service/images/core/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func UploadMultipleFiles(imgColl *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		_, id, cUser := utils.Auth(c)

		form, _ := c.MultipartForm()
		files := form.File["images[]"]

		pathName := "./uploads/" + cUser + "/"

		utils.CreateDir(pathName)

		for _, file := range files {

			extension := filepath.Ext(file.Filename)
			newFileName := uuid.New().String() + extension

			image := models.Image{
				UrlPhoto:  cUser + "/" + newFileName,
				CreatedAt: time.Now(),
				Owner:     id,
			}

			image.UpdatedAt = image.CreatedAt

			imgColl.InsertOne(context.TODO(), image)

			if err := c.SaveUploadedFile(file, pathName+newFileName); err != nil {

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Your files have been successfully uploaded.",
		})
	}
}
