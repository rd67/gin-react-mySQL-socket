package v1Users

import (
	"fmt"

	"github.com/rd67/gin-react-mySQL-socket/models"
	"github.com/rd67/gin-react-mySQL-socket/utils"
)

func UserTokenGenerate(user_id uint) (string, error) {

	token, err := utils.JwtGenerateToken(user_id)
	if err != nil {
		return "", err
	}

	var userToken = models.UserToken{
		UserId: user_id,
		Token:  token,
	}

	err = utils.DBConn.Save(&userToken).Error
	if err != nil {
		return "", err
	}

	return token, nil

}

func UserAuth(email string, password string) (string, error) {
	var user models.User

	err := utils.DBConn.Model(models.User{}).Where(fmt.Sprintf("email = '%s'", email)).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = utils.HashMatch(user.Password, password)
	if err != nil {
		return "", err
	}

	return utils.JwtGenerateToken(user.ID)
}
