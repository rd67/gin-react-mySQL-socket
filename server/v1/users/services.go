package v1Users

import (
	"github.com/rd67/gin-react-mySQL-socket/configs"
	"github.com/rd67/gin-react-mySQL-socket/models"
	"github.com/rd67/gin-react-mySQL-socket/pkg"
)

func UserTokenGenerate(user_id uint) (string, error) {

	token, err := pkg.JwtGenerateToken(user_id, configs.TOKEN_TYPE_AUTH)
	if err != nil {
		return "", err
	}

	var userToken = models.UserToken{
		UserId: user_id,
		Token:  token,
	}

	err = pkg.DBConn.Save(&userToken).Error
	if err != nil {
		return "", err
	}

	return token, nil

}
