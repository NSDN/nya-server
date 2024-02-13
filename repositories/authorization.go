package repositories

import (
	"context"

	"github.com/NSDN/nya-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

// 获取目标用户的验证信息
func GetTargetAuthorizationInfo(username string) (models.UserAuthorizateInfo, error) {
	collection := getUsersCollection()

	filter := bson.D{{Key: "username", Value: username}}

	var result models.UserAuthorizateInfo
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}
