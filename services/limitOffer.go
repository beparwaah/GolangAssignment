package services

import (
	"errors"
	"time"

	"github.com/beparwaah/assignmentVegapay/constants"
	"github.com/beparwaah/assignmentVegapay/daoRepo"
	"github.com/beparwaah/assignmentVegapay/models"
)

func CreateLimitOffer(offer *models.LimitOffer) error {
	if offer.NewLimit <= 0 {
		return errors.New("New limit must be greater than 0")
	}

	if offer.OfferActivationTime.IsZero() || offer.OfferActivationTime.Before(time.Now()) {
		return errors.New("Invalid offer activation time")
	}

	if offer.OfferExpiryTime.IsZero() || offer.OfferExpiryTime.Before(offer.OfferActivationTime) {
		return errors.New("Invalid offer expiry time")
	}
	// checking if account exists
	account, err := daoRepo.GetAccountByID(offer.AccountID)
	if err != nil {
		return err
	}
	if account == nil {
		return errors.New("Account not exits")
	}
	//add it in constants model
	if offer.LimitType == constants.AccountLimitType && account.AccountLimit >= offer.NewLimit {
		return errors.New("Limit should be greater than previous limit")
	}
	if offer.LimitType == constants.PerTransactionLimitType && account.PerTransactionLimit >= offer.NewLimit {
		return errors.New("Limit should be greater than previous limit")
	}
	offer.OfferActivationTime = time.Now()
	// setting offer status as PENDING
	offer.Status = "PENDING"

	err = daoRepo.CreateLimitOffer(offer)
	if err != nil {
		return err
	}
	return nil
}
