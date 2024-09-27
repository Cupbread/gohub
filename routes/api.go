package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!!",
			})
		})
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignUpController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
