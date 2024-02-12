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
	customer := []Customer{}
	query := "SELECT customer_id, name, date_of_birth, city, zip_code, status FROM customers"
	err := r.db.Select(&customer, query)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r customerRepositortyDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT customer_id, name, date_of_birth, city, zip_code, status FROM customers WHERE customer_id = $1"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
