package v1Users

import (
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/models"
)

type UserRegisterDataResponseStruct struct {
	User models.User `json:"user"`
}
type UserRegisterResponseStruct struct {
	configs.CommonResponseStruct
	Data UserRegisterDataResponseStruct `json:"data"`
}
