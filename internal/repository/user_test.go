package repository

import (
	"context"
	"regexp"
	models "test-plus/internal/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUserRepository_InsertNewUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	assert.NoError(t, err)

	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    error
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				user: &models.User{
					Username:    "Agus",
					Email:       "agus@gmail.com",
					PhoneNumber: "085748949422",
					FullName:    "Agus Sehat",
					Address:     "Jakarta",
					Dob:         "2012-09-02",
					Password:    "123456",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			wantErr: false,
			want:    nil,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO `users` (`username`,`email`,`phone_number`,`full_name`,`address`,`dob`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?)",
				)).
					WithArgs(
						args.user.Username,
						args.user.Email,
						args.user.PhoneNumber,
						args.user.FullName,
						args.user.Address,
						args.user.Dob,
						args.user.Password,
						args.user.CreatedAt,
						args.user.UpdatedAt,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				user: &models.User{
					Username:    "Agus",
					Email:       "agus@gmail.com",
					PhoneNumber: "085748949422",
					FullName:    "Agus Sehat",
					Address:     "Jakarta",
					Dob:         "2012-09-02",
					Password:    "123456",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			wantErr: true,
			want:    nil,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(regexp.QuoteMeta(
					"INSERT INTO `users` (`username`,`email`,`phone_number`,`full_name`,`address`,`dob`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?)",
				)).
					WithArgs(
						args.user.Username,
						args.user.Email,
						args.user.PhoneNumber,
						args.user.FullName,
						args.user.Address,
						args.user.Dob,
						args.user.Password,
						args.user.CreatedAt,
						args.user.UpdatedAt,
					).
					WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &UserRepository{
				DB: gormDB,
			}
			if err := r.InsertNewUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.InsertNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
