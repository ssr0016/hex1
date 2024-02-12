package service

import (
	"database/sql"

	"github.com/ssr0016/goHex/errors"
	"github.com/ssr0016/goHex/logs"
	"github.com/ssr0016/goHex/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{
		custRepo: custRepo,
	}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errors.NewUnexpectedError()
	}

	custReponses := []CustomerResponse{}
	for _, customer := range customers {
		custReponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custReponses = append(custReponses, custReponse)
	}

	return custReponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errors.NewUnexpectedError()
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
