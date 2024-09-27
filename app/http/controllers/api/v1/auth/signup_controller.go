package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"net/http"
)

type SignUpController struct {
	v1.BaseApiController
}

func (sc SignUpController) IsPhoneExist(c *gin.Context) {
	type PhoneExistRequest struct {
		Phone string `json:"phone" binding:"required"`
	}
	request := PhoneExistRequest{}

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(request.Phone)})
}
