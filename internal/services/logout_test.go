package services

import (
	"context"
	"test-plus/internal/mocks"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestLogoutService_Logout(t *testing.T) {
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
		wantErr bool
		mockFn  func()
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LogoutService{
				UserRepo: mockRepo,
			}
			if err := s.Logout(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("LogoutService.Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
