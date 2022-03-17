package domain

import (
	"github.com/ravikrs/go-rest-banking/s3/dto"
	"github.com/ravikrs/go-rest-banking/s3/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

func (c Customer) statusAsText() string {
	statusText := "inactive"
	if c.Status == "1" {
		statusText = "active"
	}
	return statusText
}

type Repository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindByID(id string) (*Customer, *errs.AppError)
	FindByStatus(status string) ([]Customer, *errs.AppError)
}
