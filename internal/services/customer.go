package services

import (
	"context"
	"fmt"
	"test-plus/internal/interfaces"
	models "test-plus/internal/model"
	"time"
)

type CustomerService struct {
	CustomerRepo interfaces.ICustomerRepo
}

func (s *CustomerService) CreateCustomer(ctx context.Context, customer *models.CustomerParam) (*models.Customer, error) {
	var (
		obj models.Customer
	)
	obj.Nik = customer.Nik
	obj.FullName = customer.FullName
	obj.LegalName = customer.LegalName
	obj.TempatLahir = customer.TempatLahir
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, customer.TanggalLahir)
	if err != nil {
		return nil, fmt.Errorf("invalid date format")
	}
	obj.TanggalLahir = parsedTime
	obj.Gaji = customer.Gaji
	obj.FotoKtp = customer.FotoKtp
	obj.FotoSelfi = customer.FotoSelfi
	data, err := s.CustomerRepo.CreateCustomer(ctx, &obj)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *CustomerService) FindByID(ctx context.Context, ID int) (models.Customer, error) {
	return s.CustomerRepo.FindByID(ctx, ID)
}
