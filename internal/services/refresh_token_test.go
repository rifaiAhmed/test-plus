package services

import (
	"context"
	"reflect"
	"test-plus/helpers"
	"test-plus/internal/mocks"
	models "test-plus/internal/model"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestRefreshTokenService_RefreshToken(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrlMock)
	type args struct {
		ctx          context.Context
		refreshToken string
		tokenClaim   helpers.ClaimToken
	}
	tests := []struct {
		name    string
		args    args
		want    models.RefreshTokenResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RefreshTokenService{
				UserRepo: mockRepo,
			}
			got, err := s.RefreshToken(tt.args.ctx, tt.args.refreshToken, tt.args.tokenClaim)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshTokenService.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefreshTokenService.RefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
