package interfaces

import (
	"context"
	models "test-plus/internal/model"
)

//go:generate mockgen -source=IUser.go -destination=../mocks/IUser_mock.go -package=mocks
type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.User) error
	GetUserbyUsername(ctx context.Context, username string) (models.User, error)
	InsertNewUserSession(ctx context.Context, session *models.UserSession) error
	DeleteUserSession(ctx context.Context, token string) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSession, error)
	UpdateTokenWByRefreshToken(ctx context.Context, token string, refreshToken string) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSession, error)
}
