package repositories

import (
	"github.com/NSDN/nya-server/context"
	"github.com/NSDN/nya-server/models"
)

type PlateRepository struct {
	context *context.AppContext
}

func NewPlateRepository(context *context.AppContext) *PlateRepository {
	return &PlateRepository{context: context}
}

/*
创建版块列表 - 数据库

	@param plates models/plates > Plate
	@returns _id[], err
*/
func (repository *PlateRepository) InitPlateList(
	plates []models.Plate,
) ([]interface{}, error) {
	result := repository.context.DB.Create(&plates)
	return nil, result.Error
}

// 获取版块列表 - 数据库
func (repository *PlateRepository) GetPlates() ([]models.Plate, error) {
	var plates []models.Plate

	// 从数据库中获取版块列表集合
	result := repository.context.DB.Find(&plates)

	if result.Error != nil {
		return nil, result.Error
	}

	return plates, nil
}
