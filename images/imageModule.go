package images

import (
	"github.com/damocles217/images_service/images/api/cache"
	"github.com/damocles217/images_service/images/api/controllers"
	"github.com/damocles217/images_service/images/api/database"
	"github.com/damocles217/images_service/images/api/guards"
	"github.com/damocles217/images_service/middlewares"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {

	collection := database.MongoStart("mongodb://localhost:27017/")
	collectionImage := database.MongoStartImg("mongodb://localhost:27017/")
	cacheUser := cache.CreateCacheServer()

	/*
		TODO Set this mode for production: gin.SetMode(gin.ReleaseMode)
	*/
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.SetTrustedProxies([]string{"192.168.1.65"})

	r.Use(middlewares.CORSMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.MaxMultipartMemory = 10 << 20 // 8 MiB

	// ! TODO Connect to other microservice for auth guard, because using two collections
	// ! is too many slowly the api
	images := r.Group("/images")
	{
		images.GET("/getmyownimages/:page", controllers.GetMyImages(collectionImage, cacheUser))
		images.GET("/getimage/:user/:id", controllers.GetFile)

		images.DELETE("/deleteimage", controllers.DeleteImage(collectionImage, cacheUser))

		images.POST("/upload", guards.AuthGuard(), controllers.UploadFile(collectionImage, cacheUser))
		images.POST("/profile", guards.AuthGuard(), controllers.SetProfileImg(collection))
		images.POST("/uploadmultiple", guards.AuthGuard(), controllers.UploadMultipleFiles(collectionImage))
	}

	return r
}
