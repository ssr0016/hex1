package service

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ssr0016/goHex/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{
		custRepo: custRepo,
	}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custReponses := []CustomerResponse{}
	for _, cust := range customers {
		custReponse := CustomerResponse{
			CustomerID: cust.CustomerID,
			Name:       cust.Name,
			Status:     cust.Status,
		}
		custReponses = append(custReponses, custReponse)
	}

	return custReponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		log.Println(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
