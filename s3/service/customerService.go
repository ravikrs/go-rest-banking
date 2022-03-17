package service

import (
	"github.com/ravikrs/go-rest-banking/s3/domain"
	"github.com/ravikrs/go-rest-banking/s3/dto"
	"github.com/ravikrs/go-rest-banking/s3/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerByID(id string) (*dto.CustomerResponse, *errs.AppError)
	GetAllCustomerByStatus(status string) ([]dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.Repository
}

func (service DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errs.AppError) {
	customers, appError := service.repo.FindAll()
	if appError != nil {
		return nil, appError
	}
	customersResponse := make([]dto.CustomerResponse, 0)
	for _, v := range customers {
		customersResponse = append(customersResponse, v.ToDTO())
	}
	return customersResponse, nil

}

func (service DefaultCustomerService) GetCustomerByID(id string) (*dto.CustomerResponse, *errs.AppError) {
	byID, appError := service.repo.FindByID(id)
	if appError != nil {
		return nil, appError
	}
	toDTO := byID.ToDTO()
	return &toDTO, nil
}

func (service DefaultCustomerService) GetAllCustomerByStatus(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, appError := service.repo.FindByStatus(status)
	if appError != nil {
		return nil, appError
	}
	customersResponse := make([]dto.CustomerResponse, 0)
	for _, v := range customers {
		customersResponse = append(customersResponse, v.ToDTO())
	}
	return customersResponse, nil

}

func NewDefaultCustomerService(repo domain.Repository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
