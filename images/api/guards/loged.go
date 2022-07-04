package guards

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/damocles217/images_service/images/api/utils"
	"github.com/gin-gonic/gin"
)

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := &http.Client{}

		code_auth, _, _ := utils.Auth(c)

		req, _ := http.NewRequest("GET", "http://192.168.1.65:3000/user", nil)
		c_user := &http.Cookie{
			Name:  "c_user",
			Value: code_auth,
		}

		tUser, _ := c.Cookie("t_user")

		t_user := &http.Cookie{
			Name:  "t_user",
			Value: tUser,
		}

		req.AddCookie(c_user)
		req.AddCookie(t_user)
		res, _ := client.Do(req)
		body, _ := ioutil.ReadAll(res.Body)

		bodyBool, _ := strconv.ParseBool(string(body))

		if bodyBool {
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
