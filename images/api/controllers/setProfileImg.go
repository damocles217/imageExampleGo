package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/damocles217/images_service/images/api/utils"
	"github.com/damocles217/images_service/images/core/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetProfileImg(userColl *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		var photo models.Image
		user := models.User{}

		if err := c.BindJSON(&photo); err != nil {
			return
		}

		_, id, _ := utils.Auth(c)

		idObj, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			log.Fatal(err.Error())
		}

		filter := bson.D{primitive.E{Key: "_id", Value: idObj}}

		update := bson.D{primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "url_photo", Value: photo.UrlPhoto},
		}}}

		userColl.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)

		c.JSON(http.StatusOK, gin.H{
			"url_photo": user.UrlPhoto,
		})
	}
}
