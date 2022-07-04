package controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/damocles217/images_service/images/api/utils"
	"github.com/damocles217/images_service/images/core/schemas"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteImage(imgColl *mongo.Collection, cacheUser *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteImage schemas.DeleteImg

		_, id, _ := utils.Auth(c)

		cacheUser.Del(context.TODO(), "images_"+id)

		err := c.ShouldBind(&deleteImage)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Not valid form",
			})
			return
		}

		if deleteImage.Validate != true {
			c.JSON(http.StatusLocked, gin.H{
				"message": "Need to confirm if you want to delete",
			})
			return
		}
		err = os.Remove("./uploads/" + deleteImage.UrlPhoto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "File not found",
			})
		}

		filter := bson.D{primitive.E{Key: "url_photo", Value: deleteImage.UrlPhoto}, primitive.E{Key: "owner", Value: id}}
		data := imgColl.FindOneAndDelete(context.TODO(), filter)

		c.JSON(http.StatusAccepted, data)

	}
}
