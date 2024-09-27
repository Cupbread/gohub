package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"
)

type SignUpController struct {
	v1.BaseApiController
}

func (sc SignUpController) IsPhoneExist(c *gin.Context) {
	request := requests.SignupPhoneExistRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	errs := requests.ValidateSignupPhoneExist(&request, c)

	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(request.Phone)})
}
