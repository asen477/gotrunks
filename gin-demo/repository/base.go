package repository

import (
	"gotrunks/gin-demo/core/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type baseRepository struct {
	db *gorm.DB
}

func (b *baseRepository) New(ctx *gin.Context) {
	b.db = database.DB(ctx)
}
