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
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(request.Phone)})
}

func (sc SignUpController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{"exist": user.IsEmailExist(request.Email)})
}
