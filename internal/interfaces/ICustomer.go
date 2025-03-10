package interfaces

import (
	"context"
	models "test-plus/internal/model"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=ICustomer.go -destination=../mocks/ICustomer_mock.go -package=mocks
type ICustomerRepo interface {
	CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	FindByID(ctx context.Context, ID int) (models.Customer, error)
}

type ICustomerService interface {
	CreateCustomer(ctx context.Context, customer *models.CustomerParam) (*models.Customer, error)
	FindByID(ctx context.Context, ID int) (models.Customer, error)
}

type ICustomerAPI interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
}
