package repository

import (
	"context"
	"errors"
	"fmt"
	models "test-plus/internal/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func (r *TransactionRepo) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	tx := r.DB.Begin()

	if transaction.JumlahBulan > 4 || transaction.JumlahBulan == 0 {
		return nil, errors.New("tenor tidak boleh lebih dari 4 bulan atau di bawah 1 bulan")
	}

	code, err := GenerateCode(tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// cek limit
	errUpdateLimit := CekCreditLimit(tx, int(transaction.CustomerID), transaction.Otr, int(transaction.JumlahBulan))
	if errUpdateLimit != nil {
		tx.Rollback()
		return nil, err
	}
	// update limit
	transaction.NomorKontrak = code
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("gagal commit transaksi: %v", err)
	}
	return transaction, nil
}

func GenerateCode(tx *gorm.DB) (string, error) {
	var count int64
	if err := tx.Model(&models.Transaction{}).Count(&count).Error; err != nil {
		return "", err
	}
	currentDate := time.Now().Format("060102")

	transactionNumber := count + 1

	code := fmt.Sprintf("TR-%s-%05d", currentDate, transactionNumber)
	return code, nil
}

func CekCreditLimit(tx *gorm.DB, customerId int, amount float64, tenor int) error {
	var obj models.CreditLimit

	// Locking menggunakan FOR UPDATE
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("customer_id = ?", customerId).
		First(&obj).Error; err != nil {
		return err
	}

	// Cek limit berdasarkan tenor
	switch tenor {
	case 1:
		if amount > obj.Tenor1 {
			return errors.New("limit tenor 1 tidak mencukupi")
		}
		obj.Tenor1 -= amount
	case 2:
		if amount > obj.Tenor2 {
			return errors.New("limit tenor 2 tidak mencukupi")
		}
		obj.Tenor2 -= amount
	case 3:
		if amount > obj.Tenor3 {
			return errors.New("limit tenor 3 tidak mencukupi")
		}
		obj.Tenor3 -= amount
	case 4:
		if amount > obj.Tenor4 {
			return errors.New("limit tenor 4 tidak mencukupi")
		}
		obj.Tenor4 -= amount
	default:
		return errors.New("tenor tidak valid, harus 1, 2, 3, atau 4")
	}

	// Simpan perubahan
	if err := tx.Save(&obj).Error; err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepo) FindByTranscID(ctx context.Context, ID int) (models.Transaction, error) {
	var (
		resp = models.Transaction{}
	)
	if err := r.DB.Preload("Customer").Where("id = ?", ID).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}
