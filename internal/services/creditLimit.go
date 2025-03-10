package services

import (
	"context"
	"test-plus/internal/interfaces"
	models "test-plus/internal/model"
)

type CreditLimitService struct {
	CreditLimitRepo interfaces.ICreditLimitRepo
}

func (s *CreditLimitService) CreateCreditLimit(ctx context.Context, creditLimit *models.CreditLimit) (*models.CreditLimit, error) {
	data, err := s.CreditLimitRepo.CreateCreditLimit(ctx, creditLimit)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *CreditLimitService) FindLimitByID(ctx context.Context, ID int) (models.CreditLimit, error) {
	return s.CreditLimitRepo.FindByID(ctx, ID)
}
