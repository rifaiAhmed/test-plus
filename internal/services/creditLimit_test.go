package services

import (
	"context"
	"errors"
	"reflect"
	"test-plus/internal/mocks"
	models "test-plus/internal/model"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func TestCreditLimitService_CreateCreditLimit(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	now := time.Now()

	mockRepo := mocks.NewMockICreditLimitRepo(ctrlMock)
	type args struct {
		ctx         context.Context
		creditLimit *models.CreditLimit
	}
	tests := []struct {
		name    string
		args    args
		want    *models.CreditLimit
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				creditLimit: &models.CreditLimit{
					CustomerID: 2,
					Tenor1:     100000,
					Tenor2:     200000,
					Tenor3:     300000,
					Tenor4:     400000,
				},
			},
			want: &models.CreditLimit{
				ID:         1,
				CustomerID: 2,
				Tenor1:     100000,
				Tenor2:     200000,
				Tenor3:     300000,
				Tenor4:     400000,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().CreateCreditLimit(args.ctx, gomock.Any()).Return(&models.CreditLimit{
					ID:         1,
					CustomerID: 2,
					Tenor1:     100000,
					Tenor2:     200000,
					Tenor3:     300000,
					Tenor4:     400000,
					CreatedAt:  now,
					UpdatedAt:  now,
				}, nil)
			},
		},
		{
			name: "error case",
			args: args{
				ctx: context.Background(),
				creditLimit: &models.CreditLimit{
					CustomerID: 2,
					Tenor1:     100000,
					Tenor2:     200000,
					Tenor3:     300000,
					Tenor4:     400000,
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().CreateCreditLimit(args.ctx, gomock.Any()).Return(nil, errors.New("database error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}
			s := &CreditLimitService{
				CreditLimitRepo: mockRepo,
			}
			got, err := s.CreateCreditLimit(tt.args.ctx, tt.args.creditLimit)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreditLimitService.CreateCreditLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreditLimitService.CreateCreditLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreditLimitService_FindLimitByID(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	now := time.Now()

	mockRepo := mocks.NewMockICreditLimitRepo(ctrlMock)
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    models.CreditLimit
		wantErr bool
		mockFn  func()
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			wantErr: false,
			want: models.CreditLimit{
				ID:        1,
				Tenor1:    100000,
				Tenor2:    200000,
				Tenor3:    300000,
				Tenor4:    400000,
				CreatedAt: now,
				UpdatedAt: now,
			},
			mockFn: func() {
				mockRepo.EXPECT().FindByID(gomock.Any(), 1).Return(models.CreditLimit{
					ID:        1,
					Tenor1:    100000,
					Tenor2:    200000,
					Tenor3:    300000,
					Tenor4:    400000,
					CreatedAt: now,
					UpdatedAt: now,
				}, nil)
			},
		},
		{
			name: "error limit not found",
			args: args{
				ctx: context.Background(),
				ID:  999,
			},
			wantErr: true,
			want:    models.CreditLimit{},
			mockFn: func() {
				mockRepo.EXPECT().
					FindByID(gomock.Any(), 999).
					Return(models.CreditLimit{}, errors.New("limit not found"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CreditLimitService{
				CreditLimitRepo: mockRepo,
			}
			tt.mockFn()
			got, err := s.FindLimitByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreditLimitService.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreditLimitService.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
