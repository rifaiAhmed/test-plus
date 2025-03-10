package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=ILogout.go -destination=../mocks/ILogout_mock.go -package=mocks
type ILogoutService interface {
	Logout(ctx context.Context, token string) error
}

type ILogoutHandler interface {
	Logout(*gin.Context)
}
