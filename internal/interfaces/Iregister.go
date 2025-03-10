package interfaces

import (
	"context"
	models "test-plus/internal/model"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=Iregister.go -destination=../mocks/Iregister_mock.go -package=mocks
type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
type IRegisterHandler interface {
	Register(*gin.Context)
}
