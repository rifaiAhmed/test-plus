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

func TestCustomerService_CreateCustomer(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	now := time.Now()

	mockRepo := mocks.NewMockICustomerRepo(ctrlMock)

	type args struct {
		ctx      context.Context
		customer *models.CustomerParam
	}

	tests := []struct {
		name    string
		args    args
		want    *models.Customer
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				customer: &models.CustomerParam{
					Nik:          "1234567890123456",
					FullName:     "Budi",
					LegalName:    "Budi Santoso",
					TempatLahir:  "Jakarta",
					TanggalLahir: "2002-11-09",
					Gaji:         7000000,
					FotoKtp:      "ktp.jpg",
					FotoSelfi:    "selfie.jpg",
				},
			},
			want: &models.Customer{
				ID:          1,
				Nik:         "1234567890123456",
				FullName:    "Budi",
				LegalName:   "Budi Santoso",
				TempatLahir: "Jakarta",
				TanggalLahir: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2002-11-09")
					return t
				}(),
				Gaji:      7000000,
				FotoKtp:   "ktp.jpg",
				FotoSelfi: "selfie.jpg",
				CreatedAt: now,
				UpdatedAt: now,
			},
			wantErr: false,
			mockFn: func(args args) {
				parsedTime, _ := time.Parse("2006-01-02", args.customer.TanggalLahir)
				mockRepo.EXPECT().CreateCustomer(args.ctx, gomock.Any()).Return(&models.Customer{
					ID:           1,
					Nik:          args.customer.Nik,
					FullName:     args.customer.FullName,
					LegalName:    args.customer.LegalName,
					TempatLahir:  args.customer.TempatLahir,
					TanggalLahir: parsedTime,
					Gaji:         args.customer.Gaji,
					FotoKtp:      args.customer.FotoKtp,
					FotoSelfi:    args.customer.FotoSelfi,
					CreatedAt:    now,
					UpdatedAt:    now,
				}, nil)
			},
		},
		{
			name: "error parsing date",
			args: args{
				ctx: context.Background(),
				customer: &models.CustomerParam{
					Nik:          "1234567890123456",
					FullName:     "Budi",
					LegalName:    "Budi Santoso",
					TempatLahir:  "Jakarta",
					TanggalLahir: "not-valid",
					Gaji:         7000000,
					FotoKtp:      "ktp.jpg",
					FotoSelfi:    "selfie.jpg",
				},
			},
			wantErr: true,
			want:    nil,
			mockFn: func(args args) {
				layout := "2006-01-02"
				_, err := time.Parse(layout, args.customer.TanggalLahir)
				if err != nil {
					return
				}
			},
		},
		{
			name: "error creating customer",
			args: args{
				ctx: context.Background(),
				customer: &models.CustomerParam{
					Nik:          "1234567890123456",
					FullName:     "Budi",
					LegalName:    "Budi Santoso",
					TempatLahir:  "Jakarta",
					TanggalLahir: "2002-11-09",
					Gaji:         7000000,
					FotoKtp:      "ktp.jpg",
					FotoSelfi:    "selfie.jpg",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().CreateCustomer(args.ctx, gomock.Any()).Return(nil, errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}

			s := &CustomerService{
				CustomerRepo: mockRepo,
			}
			got, err := s.CreateCustomer(tt.args.ctx, tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerService.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerRepo.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerService_FindByID(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	now := time.Now()

	mockRepo := mocks.NewMockICustomerRepo(ctrlMock)

	type args struct {
		ctx context.Context
		ID  int
	}

	tests := []struct {
		name    string
		args    args
		want    models.Customer
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
			want: models.Customer{
				ID:           1,
				Nik:          "1234567890123456",
				FullName:     "Budi",
				LegalName:    "Budi Santoso",
				TempatLahir:  "Jakarta",
				TanggalLahir: now,
				Gaji:         7000000,
				FotoKtp:      "ktp.jpg",
				FotoSelfi:    "selfie.jpg",
			},
			mockFn: func() {
				mockRepo.EXPECT().
					FindByID(gomock.Any(), 1).
					Return(models.Customer{
						ID:           1,
						Nik:          "1234567890123456",
						FullName:     "Budi",
						LegalName:    "Budi Santoso",
						TempatLahir:  "Jakarta",
						TanggalLahir: now,
						Gaji:         7000000,
						FotoKtp:      "ktp.jpg",
						FotoSelfi:    "selfie.jpg",
					}, nil)
			},
		},
		{
			name: "error customer not found",
			args: args{
				ctx: context.Background(),
				ID:  999,
			},
			wantErr: true,
			want:    models.Customer{},
			mockFn: func() {
				mockRepo.EXPECT().
					FindByID(gomock.Any(), 999).
					Return(models.Customer{}, errors.New("customer not found"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CustomerService{
				CustomerRepo: mockRepo,
			}

			tt.mockFn()

			got, err := s.FindByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerService.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerService.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
