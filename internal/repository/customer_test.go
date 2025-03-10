package repository_test

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	models "test-plus/internal/model"
	"test-plus/internal/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCustomerRepo_CreateCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	assert.NoError(t, err)

	type args struct {
		ctx   context.Context
		model *models.Customer
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
				model: &models.Customer{
					Nik:          "1234567890123456",
					FullName:     "Budi",
					LegalName:    "Budi Santoso",
					TempatLahir:  "Jakarta",
					TanggalLahir: time.Now(),
					Gaji:         70000000,
					FotoKtp:      "ktp.jpg",
					FotoSelfi:    "selfie.jpg",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO `customers` (`nik`,`full_name`,`legal_name`,`tempat_lahir`,`tanggal_lahir`,`gaji`,`foto_ktp`,`foto_selfi`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?)",
				)).
					WithArgs(
						args.model.Nik,
						args.model.FullName,
						args.model.LegalName,
						args.model.TempatLahir,
						args.model.TanggalLahir,
						args.model.Gaji,
						args.model.FotoKtp,
						args.model.FotoSelfi,
						args.model.CreatedAt,
						args.model.UpdatedAt,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			want: &models.Customer{
				ID:           1,
				Nik:          "1234567890123456",
				FullName:     "Budi",
				LegalName:    "Budi Santoso",
				TempatLahir:  "Jakarta",
				TanggalLahir: time.Now(),
				Gaji:         70000000,
				FotoKtp:      "ktp.jpg",
				FotoSelfi:    "selfie.jpg",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
		},
		{
			name: "error case",
			args: args{
				ctx: context.Background(),
				model: &models.Customer{
					Nik:          "1234567890123456",
					FullName:     "Budi",
					LegalName:    "Budi Santoso",
					TempatLahir:  "Jakarta",
					TanggalLahir: time.Now(),
					Gaji:         70000000,
					FotoKtp:      "ktp.jpg",
					FotoSelfi:    "selfie.jpg",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO `customers` (`nik`,`full_name`,`legal_name`,`tempat_lahir`,`tanggal_lahir`,`gaji`,`foto_ktp`,`foto_selfi`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?)",
				)).
					WithArgs(
						args.model.Nik,
						args.model.FullName,
						args.model.LegalName,
						args.model.TempatLahir,
						args.model.TanggalLahir,
						args.model.Gaji,
						args.model.FotoKtp,
						args.model.FotoSelfi,
						args.model.CreatedAt,
						args.model.UpdatedAt,
					).
					WillReturnError(errors.New("insert failed"))

				mock.ExpectRollback()
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository.CustomerRepo{
				DB: gormDB,
			}
			got, err := r.CreateCustomer(tt.args.ctx, tt.args.model)

			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerRepo.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				assert.Equal(t, tt.want.ID, got.ID)
				assert.Equal(t, tt.want.Nik, got.Nik)
				assert.Equal(t, tt.want.LegalName, got.LegalName)
			} else {
				assert.Nil(t, got)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestCustomerRepo_FindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	now := time.Now()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	assert.NoError(t, err)
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    models.Customer
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: models.Customer{
				ID:           1,
				Nik:          "1234567890123456",
				FullName:     "Budi",
				LegalName:    "Budi Santoso",
				TempatLahir:  "Jakarta",
				TanggalLahir: now,
				Gaji:         70000000,
				FotoKtp:      "ktp.jpg",
				FotoSelfi:    "selfie.jpg",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			wantErr: false,
			mockFn: func(args args) {

				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `customers` WHERE id = ? ORDER BY `customers`.`id` LIMIT ?",
				)).WithArgs(args.ID, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "created_at", "updated_at", "nik", "full_name", "legal_name",
						"tempat_lahir", "tanggal_lahir", "gaji", "foto_ktp", "foto_selfi",
					}).AddRow(
						1, now, now, "1234567890123456", "Budi", "Budi Santoso",
						"Jakarta", now, 70000000, "ktp.jpg", "selfie.jpg",
					))

			},
		},
		{
			name: "case error",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want:    models.Customer{},
			wantErr: true,
			mockFn: func(args args) {

				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `customers` WHERE id = ? ORDER BY `customers`.`id` LIMIT ?",
				)).WithArgs(args.ID, 1).
					WillReturnError(assert.AnError)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository.CustomerRepo{
				DB: gormDB,
			}
			got, err := r.FindByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerRepo.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerRepo.FindByID() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
