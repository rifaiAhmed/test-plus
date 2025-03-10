package services

import (
	"context"
	"reflect"
	"test-plus/internal/mocks"
	models "test-plus/internal/model"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestLoginService_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrlMock)

	type args struct {
		ctx context.Context
		req models.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		want    models.LoginResponse
		wantErr bool
		mockFn  func()
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LoginService{
				UserRepo: mockRepo,
			}
			got, err := s.Login(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
