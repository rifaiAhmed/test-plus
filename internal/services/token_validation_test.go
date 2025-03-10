package services

import (
	"context"
	"reflect"
	"test-plus/helpers"
	"test-plus/internal/mocks"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestTokenValidationService_TokenValidation(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrlMock)
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *helpers.ClaimToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TokenValidationService{
				UserRepo: mockRepo,
			}
			got, err := s.TokenValidation(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenValidationService.TokenValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenValidationService.TokenValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}
