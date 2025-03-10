package interfaces

import (
	"context"
	"test-plus/helpers"
	models "test-plus/internal/model"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=IRefreshToken.go -destination=../mocks/IRefreshToken_mock.go -package=mocks
type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(*gin.Context)
}
