package repository

import "github.com/jmoiron/sqlx"

type customerRepositortyDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) *customerRepositortyDB {
	return &customerRepositortyDB{
		db: db,
	}
}

func (r customerRepositortyDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r customerRepositortyDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id = $1"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
