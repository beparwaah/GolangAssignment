package daoRepo

import (
	"time"

	"github.com/beparwaah/assignmentVegapay/config"
	"github.com/beparwaah/assignmentVegapay/constants"
	"github.com/beparwaah/assignmentVegapay/models"
	"gorm.io/gorm"
	// "github.com/beparwaah/assignmentVegapay/config"
)

func CreateLimitOffer(offer *models.LimitOffer) error {
	if err := config.DB.Create(offer).Error; err != nil {
		return err
	}
	return nil
}
func GetLimitOfferByID(offerID uint) (*models.LimitOffer, error) {
	var offer models.LimitOffer
	if err := config.DB.First(&offer, offerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Offer not found
		}
		return nil, err
	}
	return &offer, nil
}
func UpdateLimitOfferStatus(offer *models.LimitOffer) error {
	return config.DB.Save(offer).Error
}
func ListActiveLimitOffers(accountID int, activeDate time.Time) ([]models.LimitOffer, error) {
	var offers []models.LimitOffer
	var status = constants.PendingStatus
	err := config.DB.Where("account_id = ? AND offer_activation_time <= ? AND offer_expiry_time >= ? AND status=?", accountID, activeDate, activeDate, status).Find(&offers).Error
	if err != nil {
		return nil, err
	}
	return offers, nil
}
