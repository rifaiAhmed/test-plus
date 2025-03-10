package services

import (
	"context"
	"test-plus/helpers"
	"test-plus/internal/interfaces"

	"github.com/pkg/errors"
)

type TokenValidationService struct {
	UserRepo interfaces.IUserRepository
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	var (
		claimToken *helpers.ClaimToken
		err        error
	)

	claimToken, err = helpers.ValidateToken(ctx, token)
	if err != nil {
		return claimToken, errors.Wrap(err, "failed to validate token")
	}

	_, err = s.UserRepo.GetUserSessionByToken(ctx, token)
	if err != nil {
		return claimToken, errors.Wrap(err, "failed to get user sesion")
	}

	return claimToken, nil
}
