package services

import (
	"errors"
	"time"

	"github.com/beparwaah/assignmentVegapay/constants"
	"github.com/beparwaah/assignmentVegapay/daoRepo"
	"github.com/beparwaah/assignmentVegapay/models"
)

func ListActiveLimitOffers(accountID int, activeDate time.Time) ([]models.LimitOffer, error) {
	// validating active date
	if activeDate.IsZero() {
		return nil, errors.New("Invalid active date")
	}

	// cehcking if account exists
	account, err := daoRepo.GetAccountByID(accountID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errors.New("Account not exist")
	}

	offers, err := daoRepo.ListActiveLimitOffers(accountID, activeDate)
	if err != nil {
		return nil, err
	}
	return offers, nil
}
func GetLimitOfferByID(offerID uint) (*models.LimitOffer, error) {
	return daoRepo.GetLimitOfferByID(offerID)
}
func UpdateLimitOfferStatus(offer *models.LimitOffer, status string) error {
	if status != constants.OfferAccepted && status != constants.OfferRejected {
		return errors.New("Invalid status")
	}

	existingOffer, err := daoRepo.GetLimitOfferByID(offer.ID)
	if err != nil {
		return err
	}
	if existingOffer == nil {
		return errors.New("Offer not found")
	}

	if existingOffer.Status != "PENDING" {
		return errors.New("Cannot update offer status. Offer is not pending")
	}
	// cehcking if account exists
	account, err := daoRepo.GetAccountByID(existingOffer.AccountID)
	if err != nil {
		return err
	}
	if account == nil {
		return errors.New("Account not found")
	}
	var accountUpdate models.Account
	// updating account limits if the offer is accepted
	if status == constants.OfferAccepted {

		if existingOffer.LimitType == string(constants.AccountLimitType) {
			accountUpdate = models.Account{
				ID:                     uint(existingOffer.AccountID),
				AccountLimit:           existingOffer.NewLimit,
				AccountLimitUpdateTime: time.Now(),
				LastAccountLimit:       account.AccountLimit,
			}
		} else {
			accountUpdate = models.Account{
				ID:                            uint(existingOffer.AccountID),
				PerTransactionLimit:           existingOffer.NewLimit,
				PerTransactionLimitUpdateTime: time.Now(),
				LastPerTransactionLimit:       account.LastPerTransactionLimit,
			}
		}
		err := daoRepo.UpdateAccount(&accountUpdate)
		if err != nil {
			return err
		}
	}

	existingOffer.Status = status
	existingOffer.UpdatedAt = time.Now()

	return daoRepo.UpdateLimitOfferStatus(existingOffer)
}
