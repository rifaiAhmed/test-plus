package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CreditLimit struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID uint      `json:"customer_id" gorm:"not null;index" valid:"required"`
	Tenor1     float64   `json:"tenor_1" gorm:"type:decimal(15,2);not null" valid:"required"`
	Tenor2     float64   `json:"tenor_2" gorm:"type:decimal(15,2);not null" valid:"required"`
	Tenor3     float64   `json:"tenor_3" gorm:"type:decimal(15,2);not null" valid:"required"`
	Tenor4     float64   `json:"tenor_4" gorm:"type:decimal(15,2);not null" valid:"required"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Customer Customer `gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE"`
}

func (*CreditLimit) TableName() string {
	return "credit_limits"
}

func (l CreditLimit) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
