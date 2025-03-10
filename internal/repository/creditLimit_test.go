package repository

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	models "test-plus/internal/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreditLimitRepo_CreateCreditLimit(t *testing.T) {
	now := time.Now()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	assert.NoError(t, err)
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
					CustomerID: 1,
					Tenor1:     100000,
					Tenor2:     200000,
					Tenor3:     300000,
					Tenor4:     400000,
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
			wantErr: false,
			want: &models.CreditLimit{
				ID:         1,
				CustomerID: 1,
				Tenor1:     100000,
				Tenor2:     200000,
				Tenor3:     300000,
				Tenor4:     400000,
			},
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO `credit_limits` (`customer_id`,`tenor1`,`tenor2`,`tenor3`,`tenor4`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)",
				)).
					WithArgs(
						args.creditLimit.CustomerID,
						args.creditLimit.Tenor1,
						args.creditLimit.Tenor2,
						args.creditLimit.Tenor3,
						args.creditLimit.Tenor4,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
					).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				creditLimit: &models.CreditLimit{
					CustomerID: 1,
					Tenor1:     100000,
					Tenor2:     200000,
					Tenor3:     300000,
					Tenor4:     400000,
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
			wantErr: true,
			want:    nil,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO `credit_limits` (`customer_id`,`tenor1`,`tenor2`,`tenor3`,`tenor4`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)",
				)).
					WithArgs(
						args.creditLimit.CustomerID,
						args.creditLimit.Tenor1,
						args.creditLimit.Tenor2,
						args.creditLimit.Tenor3,
						args.creditLimit.Tenor4,
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
					).
					WillReturnError(errors.New("insert failed"))

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &CreditLimitRepo{
				DB: gormDB,
			}

			got, err := r.CreateCreditLimit(tt.args.ctx, tt.args.creditLimit)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCreditLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
				assert.Equal(t, tt.want.ID, got.ID)
				assert.Equal(t, tt.want.CustomerID, got.CustomerID)
				assert.Equal(t, tt.want.Tenor1, got.Tenor1)
				assert.Equal(t, tt.want.Tenor2, got.Tenor2)
				assert.Equal(t, tt.want.Tenor3, got.Tenor3)
				assert.Equal(t, tt.want.Tenor4, got.Tenor4)
			}

		})
	}
}

func TestCreditLimitRepo_FindByID(t *testing.T) {
	now := time.Now()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	assert.NoError(t, err)

	type args struct {
		ctx context.Context
		ID  int
	}

	tests := []struct {
		name    string
		args    args
		want    models.CreditLimit
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: models.CreditLimit{
				ID:         1,
				CustomerID: 1,
				Tenor1:     100000,
				Tenor2:     200000,
				Tenor3:     300000,
				Tenor4:     400000,
				CreatedAt:  now,
				UpdatedAt:  now,
				Customer: models.Customer{
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
			},
			wantErr: false,
			mockFn: func(args args) {

				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `credit_limits` WHERE id = ? ORDER BY `credit_limits`.`id` LIMIT ?",
				)).WithArgs(args.ID, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "created_at", "updated_at", "customer_id", "tenor1", "tenor2",
						"tenor3", "tenor4",
					}).AddRow(
						1, now, now, 1, 100000, 200000, 300000, 400000,
					))

				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `customers` WHERE `customers`.`id` = ?",
				)).WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "nik", "full_name", "legal_name", "tempat_lahir", "tanggal_lahir",
						"gaji", "foto_ktp", "foto_selfi", "created_at", "updated_at",
					}).AddRow(
						1, "1234567890123456", "Budi", "Budi Santoso", "Jakarta", now,
						70000000, "ktp.jpg", "selfie.jpg", now, now,
					))
			},
		},
		{
			name: "error: credit limit not found",
			args: args{
				ctx: context.Background(),
				ID:  999,
			},
			want:    models.CreditLimit{},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `credit_limits` WHERE id = ? ORDER BY `credit_limits`.`id` LIMIT ?",
				)).WithArgs(args.ID, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "created_at", "updated_at", "customer_id", "tenor1", "tenor2",
						"tenor3", "tenor4",
					}))
			},
		},
		{
			name: "error: database query failed",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want:    models.CreditLimit{},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `credit_limits` WHERE id = ? ORDER BY `credit_limits`.`id` LIMIT ?",
				)).WithArgs(args.ID, 1).
					WillReturnError(errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &CreditLimitRepo{
				DB: gormDB,
			}
			got, err := r.FindByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreditLimitRepo.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreditLimitRepo.FindByID() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
