package repository

import (
	"context"
	models "test-plus/internal/model"

	"gorm.io/gorm"
)

type CustomerRepo struct {
	DB *gorm.DB
}

func (r *CustomerRepo) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	if err := r.DB.Create(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepo) FindByID(ctx context.Context, ID int) (models.Customer, error) {
	var (
		resp = models.Customer{}
	)
	if err := r.DB.Where("id = ?", ID).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}
