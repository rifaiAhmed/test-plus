package repository

import (
	"gorm.io/gorm"
)

type LoginRepository struct {
	DB *gorm.DB
}
