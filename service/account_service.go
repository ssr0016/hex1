package service

import (
	"time"

	"github.com/ssr0016/goHex/errors"
	"github.com/ssr0016/goHex/logs"
	"github.com/ssr0016/goHex/repository"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{
		accRepo: accRepo,
	}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	//Validate
	// if request.Amount < 5000 {
	// 	return nil, errors.NewValidationError("amount at least 5,000")
	// }
	// if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
	// 	return nil, errors.NewValidationError("account type should be saving or checking")
	// }

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errors.NewUnexpectedError()
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &response, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errors.NewUnexpectedError()
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}
	return responses, nil
}
