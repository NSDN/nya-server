package repositories

import (
	"github.com/NSDN/nya-server/models"
	"gorm.io/gorm"
)

/*
创建版块列表 - 数据库

	@param plates models/plates > Plate
	@returns _id[], err
*/
func InitPlateList(db *gorm.DB, plates []models.Plate) ([]interface{}, error) {
	result := db.Create(&plates)
	return nil, result.Error
}

// 获取版块列表 - 数据库
func GetPlateList(db *gorm.DB) ([]models.Plate, error) {
	var plates []models.Plate

	// 从数据库中获取版块列表集合
	result := db.Find(&plates)

	if result.Error != nil {
		return nil, result.Error
	}

	return plates, nil
}
