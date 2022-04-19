package v1Users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/helpers"
	"github.com/rd67/gin-react-mySQL-socket/models"
	"github.com/rd67/gin-react-mySQL-socket/utils"
)

func UserRegister(c *gin.Context) {

	//	Validating Response
	var data struct {
		Name     string `form:"name" binding:"required,min=3,max=100"`
		Email    string `form:"email" binding:"required,email,max=100"`
		Password string `form:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBind(&data); err != nil {
		helpers.ValidationResponse(c, err)
		return
	}

	//	Email uniqueness check
	var emailCount int64
	if err := utils.DBConn.Model(models.User{}).Where(fmt.Sprintf("email = '%s'", data.Email)).Count(&emailCount).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}
	if emailCount > 0 {
		helpers.ActionFailedResponse(c, "Sorry, this email is already registered with us")
		return
	}

	user := models.User{}

	password, err := utils.HashString(data.Password)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	user.Name = data.Name
	user.Email = data.Email
	user.Password = password

	if err := utils.DBConn.Save(&user).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	response := UserRegisterResponseStruct{
		CommonResponseStruct: configs.CommonResponseStruct{
			StatusCode: http.StatusCreated,
			Message:    "User registered successfully",
		},
		Data: UserRegisterDataResponseStruct{
			User: user,
		},
	}

	c.JSON(response.StatusCode, response)
	return
}
