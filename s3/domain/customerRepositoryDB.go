package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ravikrs/go-rest-banking/s3/errs"
	"github.com/ravikrs/go-rest-banking/s3/logger"
	"time"
)

type CustomerRepositoryDB struct {
	//client *sql.DB
	client *sqlx.DB
}

func (repo CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	customers := make([]Customer, 0)
	err := repo.client.Select(&customers, findAllSql)
	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	//rows, err := repo.client.Query(findAllSql)
	//if err != nil {
	//	logger.Error("Error while querying customer table " + err.Error())
	//	return nil, errs.NewUnexpectedError(err.Error())
	//}
	//customers := make([]Customer, 0)
	//err = sqlx.StructScan(rows, &customers)
	//if err != nil {
	//	logger.Error("Error while scanning customers " + err.Error())
	//	return nil, errs.NewUnexpectedError("unexpected database error")
	//}
	//for rows.Next() {
	//	var c Customer
	//	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	//	if err != nil {
	//		logger.Error("Error while scanning customer " + err.Error())
	//		return nil, errs.NewUnexpectedError(err.Error())
	//	}
	//	customers = append(customers, c)
	//}
	return customers, nil
}

func (repo CustomerRepositoryDB) FindByID(id string) (*Customer, *errs.AppError) {
	//using plain go to show how much boilerplate -> prefer sqlx
	sqlQuery := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := repo.client.QueryRow(sqlQuery, id).Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error while finding customer " + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &c, nil
}

func (repo CustomerRepositoryDB) FindByStatus(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	sqlStatement := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
	err := repo.client.Select(&customers, sqlStatement, status)
	if err != nil {
		logger.Error("Error while scanning customer by status " + err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	//rows, err := repo.client.Query(sqlStatement, status)
	//if err != nil {
	//	logger.Error("Error while scanning customer by status " + err.Error())
	//	return nil, errs.NewUnexpectedError(err.Error())
	//}
	//for rows.Next() {
	//	var c Customer
	//	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	//	if err != nil {
	//		return nil, errs.NewUnexpectedError(err.Error())
	//	}
	//	customers = append(customers, c)
	//}
	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	//client, err := sql.Open("mysql", "testuser:testpassword@tcp(localhost:3306)/test")
	client, err := sqlx.Open("mysql", "testuser:testpassword@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
