package v1Users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/helpers"
	"github.com/rd67/gin-react-mySQL-socket/models"
	"github.com/rd67/gin-react-mySQL-socket/pkg"
)

//	User Register
func UserRegister(c *gin.Context) {

	//	Validating Response
	var data IUserRegisterInput
	if err := c.ShouldBind(&data); err != nil {
		helpers.ValidationResponse(c, err)
		return
	}

	//	Email uniqueness check
	var emailCount int64
	if err := models.DBConn.Model(models.User{}).Where(fmt.Sprintf("email = '%s'", data.Email)).Count(&emailCount).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}
	if emailCount > 0 {
		helpers.ActionFailedResponse(c, http.StatusBadGateway, USER_MESSAGES["EmailAlreadyRegistered"])
		return
	}

	user := models.User{}

	password, err := pkg.HashString(data.Password)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	user.Name = data.Name
	user.Email = data.Email
	user.Password = password

	if err := models.DBConn.Save(&user).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	var token string
	token, err = UserTokenGenerate(user.ID)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	response := IUserRegisterResponse{
		ICommonResponse: configs.ICommonResponse{
			StatusCode: http.StatusCreated,
			Message:    "User registered successfully",
		},
		Data: IUserRegisterDataResponse{
			Token: token,
			User:  user,
		},
	}

	c.JSON(response.StatusCode, response)
	return
}

//	User Login
func UserLogin(c *gin.Context) {
	var data IUserLoginInput

	err := c.ShouldBind(&data)
	if err != nil {
		helpers.ValidationResponse(c, err)
		return
	}

	var user models.User
	err = models.DBConn.Where(fmt.Sprintf("email = '%s'", data.Email)).First(&user).Error
	if err != nil {
		helpers.ActionFailedResponse(c, http.StatusBadRequest, USER_MESSAGES["AccountNotFound"])
		return
	}

	err = pkg.HashMatch(user.Password, data.Password)
	if err != nil {
		helpers.ActionFailedResponse(c, http.StatusBadRequest, USER_MESSAGES["AccountNotFound"])
	}

	var token string
	token, err = UserTokenGenerate(user.ID)
	if err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	response := IUserLoginResponse{
		ICommonResponse: configs.ICommonResponse{
			StatusCode: http.StatusOK,
			Message:    USER_MESSAGES["LoggedInSuccess"],
		},
		Data: IUserLoginDataResponse{
			Token: token,
			User:  user,
		},
	}

	c.JSON(response.StatusCode, response)
	return
}

//	Users Listing
func UsersListing(c *gin.Context) {
	var data = IUsersListing{
		Search: "",
		Limit:  configs.DEFAULT_LIMIT,
		Offset: 0,
	}

	//	Validation
	err := c.ShouldBindQuery(&data)
	if err != nil {
		helpers.ValidationResponse(c, err)
		return
	}

	var authUser = c.MustGet("authUser").(models.User)

	var count int64
	var rows []models.User

	var query = models.DBConn.Model(models.User{}).Where("id != ?", authUser.ID)

	if data.Search != "" {
		query.Where("(name LIKE '%" + data.Search + "%') OR (email LIKE '%" + data.Search + "%')")
	}

	if err := query.Count(&count).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	if err := query.Find(&rows).Error; err != nil {
		helpers.ErrorResponse(c, err)
		return
	}

	response := IUserListingResponse{
		ICommonResponse: configs.ICommonResponse{
			StatusCode: http.StatusOK,
			Message:    "Users listing",
		},
		Data: IUserListingData{
			Count: count,
			Rows:  rows,
		},
	}

	c.JSON(response.StatusCode, response)
	return
}
