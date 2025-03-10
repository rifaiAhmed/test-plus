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
	"gorm.io/gorm/logger"
)

func TestTransactionRepo_FindByTranscID(t *testing.T) {
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
		want    models.Transaction
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success - transaction found",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: models.Transaction{
				ID:            1,
				CustomerID:    1,
				NomorKontrak:  "TR2503090001",
				Otr:           1000000,
				JumlahCicilan: 100000,
				JumlahBunga:   2,
				JumlahBulan:   3,
				NamaAsset:     "Handphone",
				CreatedAt:     now,
				UpdatedAt:     now,
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
					"SELECT * FROM `transactions` WHERE id = ? ORDER BY `transactions`.`id` LIMIT ?",
				)).WithArgs(args.ID, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "created_at", "updated_at", "customer_id", "nomor_kontrak", "otr", "jumlah_cicilan", "jumlah_bunga",
						"jumlah_bulan", "nama_asset",
					}).AddRow(
						1, now, now, 1, "TR2503090001", 1000000, 100000, 2, 3, "Handphone",
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
			name: "error - transaction not found",
			args: args{
				ctx: context.Background(),
				ID:  99,
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `transactions` WHERE id = ? ORDER BY `transactions`.`id` LIMIT ?",
				)).WithArgs(args.ID, 1).
					WillReturnError(gorm.ErrRecordNotFound)
			},
		},

		{
			name: "error - database failure",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `transactions` WHERE id = ? ORDER BY `transactions`.`id` LIMIT ?",
				)).WithArgs(args.ID, 1).
					WillReturnError(errors.New("database error"))
			},
		},

		{
			name: "error - failed to load customer",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `transactions` WHERE id = ? ORDER BY `transactions`.`id` LIMIT ?",
				)).WithArgs(args.ID, 1).
					WillReturnRows(sqlmock.NewRows([]string{
						"id", "created_at", "updated_at", "customer_id", "nomor_kontrak", "otr", "jumlah_cicilan", "jumlah_bunga",
						"jumlah_bulan", "nama_asset",
					}).AddRow(
						1, now, now, 1, "TR2503090001", 1000000, 100000, 2, 3, "Handphone",
					))

				mock.ExpectQuery(regexp.QuoteMeta(
					"SELECT * FROM `customers` WHERE `customers`.`id` = ?",
				)).WithArgs(1).
					WillReturnError(errors.New("failed to load customer"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &TransactionRepo{
				DB: gormDB,
			}
			got, err := r.FindByTranscID(tt.args.ctx, tt.args.ID)

			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionRepo.FindByTranscID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionRepo.FindByTranscID() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGenerateCode(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Konfigurasi GORM dengan mock database
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	assert.NoError(t, err)

	// Waktu saat ini untuk memastikan format kode tetap sesuai
	currentDate := time.Now().Format("060102")

	tests := []struct {
		name    string
		mockFn  func()
		want    string
		wantErr bool
	}{
		{
			name: "success - generate code",
			mockFn: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `transactions`")).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(10)) // Simulasi ada 10 transaksi sebelumnya
			},
			want:    "TR-" + currentDate + "-00011", // Karena count = 10, maka nomor transaksi berikutnya 11
			wantErr: false,
		},
		{
			name: "error - count failed",
			mockFn: func() {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `transactions`")).
					WillReturnError(assert.AnError) // Simulasi error saat Count()
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn() // Jalankan mock sesuai skenario

			got, err := GenerateCode(gormDB)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateCode() = %v, want %v", got, tt.want)
			}

			// Pastikan semua mock expectations terpenuhi
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return gormDB, mock
}

func TestCekCreditLimit(t *testing.T) {
	db, mock := setupMockDB()

	tests := []struct {
		name       string
		customerId int
		amount     float64
		tenor      int
		mockQuery  func()
		wantErr    bool
	}{
		{
			name:       "success - tenor 1 cukup",
			customerId: 1,
			amount:     50000,
			tenor:      1,
			mockQuery: func() {
				// Mulai transaksi
				mock.ExpectBegin()

				// Mock SELECT FOR UPDATE
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `credit_limits` WHERE customer_id = ? ORDER BY `credit_limits`.`id` LIMIT ? FOR UPDATE")).
					WithArgs(1, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "tenor1", "tenor2", "tenor3", "tenor4"}).
						AddRow(1, 100000, 200000, 300000, 400000))

				// Mock UPDATE
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `credit_limits` SET `customer_id`=?, `tenor1`=?, `tenor2`=?, `tenor3`=?, `tenor4`=?, `created_at`=?, `updated_at`=? WHERE `id` = ?")).
					WithArgs(1, 50000, 200000, 300000, 400000, sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				// Commit transaksi
				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name:       "error - tenor 1 tidak cukup",
			customerId: 2,
			amount:     150000,
			tenor:      1,
			mockQuery: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `credit_limits` WHERE customer_id = ? ORDER BY `credit_limits`.`id` LIMIT ? FOR UPDATE")).
					WithArgs(2, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "tenor1", "tenor2", "tenor3", "tenor4"}).
						AddRow(2, 100000, 200000, 300000, 400000))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:       "error - customer tidak ditemukan",
			customerId: 3,
			amount:     50000,
			tenor:      1,
			mockQuery: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `credit_limits` WHERE customer_id = ? ORDER BY `credit_limits`.`id` LIMIT ? FOR UPDATE")).
					WithArgs(3, 1).
					WillReturnError(gorm.ErrRecordNotFound)
				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:       "error - query save gagal",
			customerId: 4,
			amount:     50000,
			tenor:      1,
			mockQuery: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `credit_limits` WHERE customer_id = ? ORDER BY `credit_limits`.`id` LIMIT ? FOR UPDATE")).
					WithArgs(4, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "tenor1", "tenor2", "tenor3", "tenor4"}).
						AddRow(4, 100000, 200000, 300000, 400000))

				mock.ExpectExec(regexp.QuoteMeta("UPDATE `credit_limits` SET `tenor1`=? WHERE `id` = ?")).
					WithArgs(50000, 4).
					WillReturnError(errors.New("failed to save"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:       "error - tenor tidak valid",
			customerId: 5,
			amount:     50000,
			tenor:      5,
			mockQuery: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `credit_limits` WHERE customer_id = ? ORDER BY `credit_limits`.`id` LIMIT ? FOR UPDATE")).
					WithArgs(5, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "tenor1", "tenor2", "tenor3", "tenor4"}).
						AddRow(5, 100000, 200000, 300000, 400000))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockQuery()
			tx := db.Begin()
			err := CekCreditLimit(tx, tt.customerId, tt.amount, tt.tenor)
			if (err != nil) != tt.wantErr {
				t.Errorf("CekCreditLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
			tx.Rollback()
		})
	}

	// Pastikan semua mock query telah dieksekusi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock expectations not met: %v", err)
	}
}
