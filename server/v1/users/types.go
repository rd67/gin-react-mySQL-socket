package v1Users

import (
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/models"
)

type IUserRegisterInput struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	Email    string `form:"email" binding:"required,email,max=100"`
	Password string `form:"password" binding:"required,min=6"`
}

type IUserRegisterDataResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}
type IUserRegisterResponse struct {
	configs.ICommonResponse
	Data IUserRegisterDataResponse `json:"data"`
}

//	User Login
type IUserLoginInput struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type IUserLoginDataResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}
type IUserLoginResponse struct {
	configs.ICommonResponse
	Data IUserLoginDataResponse `json:"data"`
}
