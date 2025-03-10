package services

import (
	"context"
	"reflect"
	"test-plus/internal/mocks"
	models "test-plus/internal/model"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestRegisterService_Register(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := mocks.NewMockIUserRepository(ctrlMock)
	type args struct {
		ctx     context.Context
		request models.User
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RegisterService{
				UserRepo: mockRepo,
			}
			got, err := s.Register(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterService.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
