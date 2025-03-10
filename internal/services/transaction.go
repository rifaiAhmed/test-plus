package services

import (
	"context"
	"test-plus/internal/interfaces"
	models "test-plus/internal/model"
)

type TransactionService struct {
	TransactionRepo interfaces.ITransactionRepo
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	data, err := s.TransactionRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *TransactionService) FindByTranscID(ctx context.Context, ID int) (models.Transaction, error) {
	return s.TransactionRepo.FindByTranscID(ctx, ID)
}
