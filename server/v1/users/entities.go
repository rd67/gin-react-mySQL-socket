package v1Users

import (
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/models"
)

type IUserRegisterDataResponse struct {
	User models.User `json:"user"`
}
type IUserRegisterResponse struct {
	configs.ICommonResponse
	Data IUserRegisterDataResponse `json:"data"`
}
