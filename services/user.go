package services

import (
	"errors"

	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/utils"
)

func GetUserInfo(uid string) (*models.User, error) {
	// FIXME: 从数据库中获取真实的用户信息
	users := []models.User{
		{
			UID:       "1001",
			Username:  "root",
			Nickname:  "管理员",
			UserGroup: "大喵玉",
			Icon:      "https://i.imgur.com/nNiq3Ph.jpg",
		},
	}

	var target *models.User

	for _, user := range users {

		if user.UID == user.UID {
			target = &user
			break
		}
	}

	if target == nil {
		return nil, errors.New(utils.Messages.AUTHORIZE_FAILED_NO_USER)
	}

	return target, nil
}
