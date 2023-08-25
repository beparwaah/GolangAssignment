package services

import (
	"errors"
	"time"

	"github.com/beparwaah/assignmentVegapay/config"
	"github.com/beparwaah/assignmentVegapay/daoRepo"
	"github.com/beparwaah/assignmentVegapay/models"
)

func CreateAccount(account *models.Account) error {
	// generating a unique customer_id
	var generatedCustomerID int
	if err := config.DB.Raw("SELECT nextval('customer_id_seq')").Scan(&generatedCustomerID).Error; err != nil {
		return err
	}

	//FALLBACK logic
	// checking if an account already exists for the generated customer_id
	existingAccount, err := daoRepo.GetAccountByCustomerID(generatedCustomerID)
	if err != nil {
		return err
	}
	if existingAccount != nil {
		return errors.New("This customer already associated with  an account")
	}

	// initializing  the account fields
	account.CustomerID = generatedCustomerID
	account.AccountLimit = 0
	account.PerTransactionLimit = 0
	account.LastAccountLimit = 0
	account.LastPerTransactionLimit = 0
	account.AccountLimitUpdateTime = time.Now()
	account.PerTransactionLimitUpdateTime = time.Now()
	account.CreatedAt = time.Now()

	//calling the database
	err = daoRepo.CreateAccount(account)
	if err != nil {
		// Handle error
		return err
	}
	return nil
}
func GetAccountByID(accountID int) (*models.Account, error) {
	// fetching details by account id
	account, err := daoRepo.GetAccountByID(accountID)
	if err != nil {
		return nil, err
	}
	return account, nil
}
