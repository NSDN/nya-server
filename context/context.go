package context

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppContext struct {
	DB        *gorm.DB
	APIRouter *gin.RouterGroup
}
