package context

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppContext struct {
	DB        *gorm.DB
	DBContext context.Context
	APIRouter *gin.RouterGroup
}
