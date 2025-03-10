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

func TestTransactionService_CreateTransaction(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	now := time.Now()

	mockRepo := mocks.NewMockITransactionRepo(ctrlMock)
	type args struct {
		ctx         context.Context
		transaction *models.Transaction
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Transaction
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				transaction: &models.Transaction{
					CustomerID:    1,
					NomorKontrak:  "TR2503090001",
					Otr:           7000000,
					JumlahCicilan: 1000000,
					JumlahBunga:   5,
					JumlahBulan:   4,
					NamaAsset:     "Motor",
				},
			},
			want: &models.Transaction{
				ID:            1,
				CustomerID:    1,
				NomorKontrak:  "TR2503090001",
				Otr:           7000000,
				JumlahCicilan: 1000000,
				JumlahBunga:   5,
				JumlahBulan:   4,
				NamaAsset:     "Motor",
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().CreateTransaction(args.ctx, gomock.Any()).Return(&models.Transaction{
					ID:            1,
					CustomerID:    1,
					NomorKontrak:  "TR2503090001",
					Otr:           7000000,
					JumlahCicilan: 1000000,
					JumlahBunga:   5,
					JumlahBulan:   4,
					NamaAsset:     "Motor",
					CreatedAt:     now,
					UpdatedAt:     now,
				}, nil)
			},
		},
		{
			name: "error case",
			args: args{
				ctx: context.Background(),
				transaction: &models.Transaction{
					CustomerID:    1,
					NomorKontrak:  "TR2503090001",
					Otr:           7000000,
					JumlahCicilan: 1000000,
					JumlahBunga:   5,
					JumlahBulan:   4,
					NamaAsset:     "Motor",
				},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().CreateTransaction(args.ctx, gomock.Any()).Return(nil, errors.New("database error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}
			s := &TransactionService{
				TransactionRepo: mockRepo,
			}
			got, err := s.CreateTransaction(tt.args.ctx, tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionService.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionService.CreateTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionService_FindByTranscID(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	now := time.Now()

	mockRepo := mocks.NewMockITransactionRepo(ctrlMock)
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    models.Transaction
		wantErr bool
		mockFn  func()
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: models.Transaction{
				ID:            1,
				CustomerID:    1,
				NomorKontrak:  "TR2503090001",
				Otr:           7000000,
				JumlahCicilan: 1000000,
				JumlahBunga:   5,
				JumlahBulan:   4,
				NamaAsset:     "Motor",
				CreatedAt:     now,
				UpdatedAt:     now,
			},
			wantErr: false,
			mockFn: func() {
				mockRepo.EXPECT().FindByTranscID(gomock.Any(), 1).Return(models.Transaction{
					ID:            1,
					CustomerID:    1,
					NomorKontrak:  "TR2503090001",
					Otr:           7000000,
					JumlahCicilan: 1000000,
					JumlahBunga:   5,
					JumlahBulan:   4,
					NamaAsset:     "Motor",
					CreatedAt:     now,
					UpdatedAt:     now,
				}, nil)
			},
		},
		{
			name: "error transaction not found",
			args: args{
				ctx: context.Background(),
				ID:  999,
			},
			want:    models.Transaction{},
			wantErr: true,
			mockFn: func() {
				mockRepo.EXPECT().
					FindByTranscID(gomock.Any(), 999).
					Return(models.Transaction{}, errors.New("limit not found"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TransactionService{
				TransactionRepo: mockRepo,
			}
			tt.mockFn()
			got, err := s.FindByTranscID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionService.FindByTranscID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionService.FindByTranscID() = %v, want %v", got, tt.want)
			}
		})
	}
}
