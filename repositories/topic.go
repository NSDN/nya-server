package repositories

import (
	"github.com/NSDN/nya-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// 获取帖子集合
func getTopicCollection() *mongo.Collection {
	// return getCollection(configs.DB_COLLECTION_TOPICS)
	return nil
}

// 获取帖子楼层集合
func getTopicFloorsCollection() *mongo.Collection {
	// return getCollection(configs.DB_COLLECTION_TOPIC_FLOORS)
	return nil
}

// 创建帖子信息 - 数据库
func CreateTopic(topic *models.Topic) (*mongo.InsertOneResult, error) {
	collection := getTopicCollection()
	return insertOneToCollection(collection, topic)
}

// 获取帖子列表 - 数据库
func GetTopics(plate string) (*[]models.TopicWidthID, error) {
	collection := getTopicCollection()

	topics, err := findDataFromCollection(
		&[]models.TopicWidthID{},
		collection,
		bson.D{{Key: "plate", Value: plate}},
	)

	if err != nil {
		return nil, err
	}

	return topics, nil
}

// 创建帖子楼层 - 数据库
func CreateTopicFloors(floors *models.TopicFloors) (*mongo.InsertOneResult, error) {
	collection := getTopicFloorsCollection()
	return insertOneToCollection(collection, floors)
}

// 获取帖子楼层列表 - 数据库
func GetTopicFloors(topicID string) (*models.TopicFloors, error) {
	collection := getTopicFloorsCollection()

	objectID, err := primitive.ObjectIDFromHex(topicID)

	if err != nil {
		return nil, err
	}

	floors, err := findOneDataFromCollection(
		&models.TopicFloors{},
		collection,
		bson.D{{Key: "_id", Value: objectID}},
	)

	if err != nil {
		return nil, err
	}

	return floors, nil
}
