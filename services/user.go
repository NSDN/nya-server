package services

import (
	"errors"

	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/utils"
)

func GetUserInfo(uid string) (*models.UserPublicInfo, error) {
	// users, err := repositories.GetUserList()

	// if err != nil {
	// 	return nil, err
	// }

	var target *models.UserPublicInfo

	// for _, user := range *users {
	//
	// 	if user.UID == user.UID {
	// 		target = &user
	// 		break
	// 	}
	// }

	if target == nil {
		return nil, errors.New(utils.Messages.AUTHORIZE_FAILED_NO_USER)
	}

	return target, nil
}
