package interfaces

import (
	"context"
	models "test-plus/internal/model"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=ICreditLimit.go -destination=../mocks/ICreditLimit_mock.go -package=mocks
type ICreditLimitRepo interface {
	CreateCreditLimit(ctx context.Context, creditLimit *models.CreditLimit) (*models.CreditLimit, error)
	FindByID(ctx context.Context, ID int) (models.CreditLimit, error)
}

type ICreditLimitService interface {
	CreateCreditLimit(ctx context.Context, CreditLimit *models.CreditLimit) (*models.CreditLimit, error)
	FindLimitByID(ctx context.Context, ID int) (models.CreditLimit, error)
}

type ICreditLimitAPI interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
}
