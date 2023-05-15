package services

import (
	"encoding/base64"
	"errors"
	"log"

	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/utils"
)

var FACK_SALT = "qwer1234"
var FACK_PASSWORD = "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"

// 从数据库中获取用户
func GetUserInfo(loginInfo models.LoginInfo) (*models.User, error) {
	// FIXME: 实装从数据库中获取用户列表
	base64Password, base64Salt := getFackPassword()

	if base64Password == "" || base64Salt == "" {
		return nil, errors.New("生成假数据失败")
	}

	users := []models.User{
		{
			UID:      "1001",
			Username: "root",
			Password: base64Password,
			Salt:     base64Salt,
		},
	}

	var target *models.User

	for _, user := range users {

		if user.Username == loginInfo.Username {
			target = &user
			break
		}
	}

	if target == nil {
		return nil, errors.New(utils.Messages.AUTHORIZE_FAILED_NO_USER)
	}

	return target, nil
}

// 生成假密码
func getFackPassword() (string, string) {
	base64Salt := base64.StdEncoding.EncodeToString([]byte(FACK_SALT))
	hashed, err := HashPassword(FACK_PASSWORD, base64Salt)

	if err != nil {
		log.Println(err)
		return "", ""
	}

	base64Password := base64.StdEncoding.EncodeToString(hashed)

	return base64Password, base64Salt
}
