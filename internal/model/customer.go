package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Customer struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Nik          string    `json:"nik" gorm:"type:varchar(16);unique;" valid:"required"`
	FullName     string    `json:"full_name" gorm:"type:varchar(50);" valid:"required"`
	LegalName    string    `json:"legal_name" gorm:"type:varchar(50);" valid:"required"`
	TempatLahir  string    `json:"tempat_lahir" gorm:"type:varchar(50);" valid:"required"`
	TanggalLahir time.Time `json:"tanggal_lahir" gorm:"type:date;" valid:"required"`
	Gaji         float64   `json:"gaji" gorm:"type:decimal(15,2);" valid:"required"`
	FotoKtp      string    `json:"foto_ktp" gorm:"type:varchar(255);" valid:"required"`
	FotoSelfi    string    `json:"foto_selfi" gorm:"type:varchar(255);" valid:"required"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (*Customer) TableName() string {
	return "customers"
}

func (l Customer) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
