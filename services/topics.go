package services

import (
	"strconv"

	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/dto"
	"github.com/NSDN/nya-server/models"
	"github.com/NSDN/nya-server/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopicService struct {
	context    *context.AppContext
	repository *repositories.TopicRepository
}

func NewTopicService(context *context.AppContext) *TopicService {
	return &TopicService{
		context:    context,
		repository: repositories.NewTopicRepository(context),
	}
}

// 获取帖子列表 - 服务
func (service *TopicService) GetTopics(plateID string) ([]dto.TopicListItemDTO, error) {
	topics, err := service.repository.GetTopics(plateID)

	if err != nil {
		return nil, err
	}

	result := make([]dto.TopicListItemDTO, 0, len(topics))

	for _, topic := range topics {
		// TODO: 将 `strconv.FormatInt` 更改为真的编码方法。
		result = append(result, dto.TopicListItemDTO{
			ID:            strconv.FormatInt(topic.ID, 10),
			Title:         topic.Title,
			ThumbnailLink: topic.ThumbnailLink,
			UpdatedAt:     topic.UpdatedAt,
		})
	}

	return result, nil
}

// 创建帖子 - 服务
//
// 将帖子信息和楼层分为两部分存入数据库。
// 楼层使用创建帖子信息时由数据库自增生成的 ID 作为唯一键。
func CreateTopic(request *models.NewTopicRequestData) (bool, error) {
	topic := models.Topic{
		Author:  request.Author,
		PlateID: request.Plate,
		Title:   request.Title,
		Tag:     request.Tag,
	}

	insertResult, err := repositories.CreateTopic(&topic)

	if err != nil {
		return false, err
	}

	_, err = CreateTopicFloors(request, insertResult.InsertedID.(primitive.ObjectID))

	if err != nil {
		return false, err
	}

	return true, nil
}

// 创建帖子楼层 - 服务
func CreateTopicFloors(
	request *models.NewTopicRequestData,
	id primitive.ObjectID,
) (*mongo.InsertOneResult, error) {
	floors := models.TopicFloors{
		TopicID: id,
		List:    []models.FloorItem{},
	}

	floors.List = append(floors.List, models.FloorItem{
		Level:        1,
		Author:       request.Author,
		CreationDate: request.CreationDate,
		BodyType:     request.BodyType,
		Body:         request.Body,
	})

	return repositories.CreateTopicFloors(&floors)
}

// 获取帖子列表 - 服务
//
// 从数据库中找出所有帖子，然后根据传入的版块路由名来过筛选出当前版块的帖子列表。
func GetTopicsOld(plate string) (*[]models.TopicWidthID, error) {
	topics, err := repositories.GetTopicsOld(plate)

	if err != nil {
		return nil, err
	}

	return topics, nil
}

// 获取帖子列表 - 服务
//
// 从数据库中找出所有帖子，然后根据传入的版块路由名来过筛选出当前版块的帖子列表。
func GetTopicFloors(topicID string) (*models.TopicFloors, error) {
	floors, err := repositories.GetTopicFloors(topicID)

	if err != nil {
		return nil, err
	}

	return floors, nil
}
