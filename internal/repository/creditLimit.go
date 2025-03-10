package repository

import (
	"context"
	models "test-plus/internal/model"

	"gorm.io/gorm"
)

type CreditLimitRepo struct {
	DB *gorm.DB
}

func (r *CreditLimitRepo) CreateCreditLimit(ctx context.Context, creditLimit *models.CreditLimit) (*models.CreditLimit, error) {
	result := r.DB.Create(creditLimit)
	if result.Error != nil {
		return nil, result.Error
	}
	return creditLimit, nil
}

func (r *CreditLimitRepo) FindByID(ctx context.Context, ID int) (models.CreditLimit, error) {
	var (
		resp = models.CreditLimit{}
	)
	if err := r.DB.Preload("Customer").Where("id = ?", ID).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}
