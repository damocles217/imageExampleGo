package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(c *gin.Context) (codeAuth string, idObj string, codeAuthUser string) {
	claims := jwt.MapClaims{
		"_id":       "",
		"code_auth": "",
	}

	cUser, err := c.Cookie("c_user")
	var code_auth string
	var id string

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No estas registrado",
		})
	}

	tUser, err := c.Cookie("t_user")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No estas registrado",
		})
	}

	_, err = jwt.ParseWithClaims(tUser, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("@f7LMw&F}$aPa/n`_c3&jkL*:V?Qf"), nil
	})

	for key, val := range claims {
		if key == "code_auth" {
			code_auth = val.(string)
		}
		if key == "_id" {
			id = val.(string)
		}
	}

	return code_auth, id, cUser
}
