package interfaces

import (
	"context"
	models "test-plus/internal/model"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=ILogin.go -destination=../mocks/ILogin_mock.go -package=mocks
type ILoginService interface {
	Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error)
}

type ILoginHandler interface {
	Login(c *gin.Context)
}
