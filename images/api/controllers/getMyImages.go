package controllers

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/damocles217/images_service/images/api/utils"
	"github.com/damocles217/images_service/images/core/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMyImages(imgColl *mongo.Collection, cacheUser *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var images []models.Image

		page := c.Param("page")
		pageInt, err := strconv.ParseInt(page, 10, 64)

		_, id, _ := utils.Auth(c)

		arr, err := cacheUser.Get(context.TODO(), "images_"+id+"_"+strconv.FormatInt(pageInt, 10)).Result()

		if arr != "" {
			err = json.Unmarshal([]byte(arr), &images)
			c.JSON(200, images)
			return
		}

		filter := bson.D{primitive.E{Key: "owner", Value: id}}
		// * value set to 9

		sort := bson.D{primitive.E{Key: "createdAt", Value: -1}}
		options := options.Find().SetSort(sort).SetLimit(10 * (pageInt + 1)).SetSkip(10 * pageInt)

		imgs, err := imgColl.Find(context.TODO(), filter, options)

		if err != nil {
			log.Fatal(err.Error())
		}

		for imgs.Next(context.TODO()) {
			var image models.Image
			err := imgs.Decode(&image)
			if err != nil {
				log.Fatal(err.Error())
			}

			images = append(images, image)
		}

		out, err := json.Marshal(images)
		if err != nil {
			print(err.Error())
		}

		cacheErr := cacheUser.Set(context.TODO(), "images_"+id+"_"+strconv.FormatInt(pageInt, 10), string(out), 60*time.Second).Err()
		if cacheErr != nil {
			print(err.Error())
		}

		c.JSON(200, images)
		return
	}

}
