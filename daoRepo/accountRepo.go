package daoRepo

import (
	"github.com/beparwaah/assignmentVegapay/config"
	"github.com/beparwaah/assignmentVegapay/models"
	"gorm.io/gorm"
)

func CreateAccount(account *models.Account) error {
	if err := config.DB.Create(account).Error; err != nil {
		return err
	}
	return nil
}
func GetAccountByID(accountID int) (*models.Account, error) {
	var account models.Account
	if err := config.DB.Where("id = ?", accountID).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Account not found
		}
		return nil, err
	}

	return &account, nil
}
func UpdateAccount(account *models.Account) error {
	if err := config.DB.Save(account).Error; err != nil {
		return err
	}
	return nil
}
func GetAccountByCustomerID(customerID int) (*models.Account, error) {
	var account models.Account
	if err := config.DB.Where("customer_id = ?", customerID).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Account not found
		}
		return nil, err
	}
	return &account, nil
}
