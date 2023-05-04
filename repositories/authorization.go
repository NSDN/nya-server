package repositories

import (
	"github.com/NSDN/nya-server/models"
)

// 获取目标用户的验证信息
func GetTargetAuthorizationInfo(username string) (models.UserAuthorizateInfo, error) {
	// collection := getUsersCollection()

	// filter := bson.D{{Key: "username", Value: username}}

	var result models.UserAuthorizateInfo
	// err := collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, nil
}
