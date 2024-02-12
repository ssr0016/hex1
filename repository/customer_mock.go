package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{
			CustomerID:  1,
			Name:        "Roar",
			DateOfBirth: "2000-01-01",
			City:        "Muntinlupa",
			ZipCode:     "12345",
			Status:      1,
		},
		{
			CustomerID:  2,
			Name:        "Boo",
			DateOfBirth: "2001-01-01",
			City:        "Manila",
			ZipCode:     "12345",
			Status:      1,
		},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
