package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID    uint      `json:"customer_id" gorm:"not null;index" valid:"required"`
	NomorKontrak  string    `json:"nomor_kontrak" gorm:"type:varchar(50);unique;not null" valid:"required"`
	Otr           float64   `json:"otr" gorm:"type:decimal(15,2);not null" valid:"required"`
	JumlahCicilan float64   `json:"jumlah_cicilan" gorm:"type:decimal(15,2);not null" valid:"required"`
	JumlahBunga   float64   `json:"jumlah_bunga" gorm:"type:decimal(15,2);not null" valid:"required"`
	JumlahBulan   float64   `json:"jumlah_bulan" gorm:"type:decimal(15,2);not null" valid:"required"`
	NamaAsset     string    `json:"nama_asset" gorm:"type:varchar(100);not null" valid:"required"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Customer Customer `gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE"`
}

func (*Transaction) TableName() string {
	return "transactions"
}

func (l Transaction) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
