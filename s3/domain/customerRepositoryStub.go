package domain

import (
	"github.com/ravikrs/go-rest-banking/s3/errs"
)

type CustomerRepositoryStub struct {
	customers []Customer
}

func (repo CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return repo.customers, nil
}
func (repo CustomerRepositoryStub) FindByID(id string) (*Customer, *errs.AppError) {
	for _, v := range repo.customers {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, errs.NewNotFoundError("customer does not exists")
}

func (repo CustomerRepositoryStub) FindByStatus(status string) ([]Customer, *errs.AppError) {
	c := make([]Customer, 0)
	for _, v := range repo.customers {
		if v.Status == status {
			c = append(c, v)
		}
	}
	return c, nil

}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "R", "Delhi", "11001", "01.01.2020", "active"},
		{"1", "K", "Mumbai", "11002", "02.01.2020", "inactive"},
		{"1", "S", "Kolkata", "11003", "03.01.2020", "active"},
	}
	return CustomerRepositoryStub{customers}
}
