package interfaces

import (
	"context"
	models "test-plus/internal/model"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=ITransaction.go -destination=../mocks/ITransaction_mock.go -package=mocks
type ITransactionRepo interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	FindByTranscID(ctx context.Context, ID int) (models.Transaction, error)
}

type ITransactionService interface {
	CreateTransaction(ctx context.Context, Transaction *models.Transaction) (*models.Transaction, error)
	FindByTranscID(ctx context.Context, ID int) (models.Transaction, error)
}

type ITransactionAPI interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
}
