package models

import (
	"time"
)

type Account struct {
	ID                            uint      `gorm:"primaryKey;not null" json:"id"`
	CustomerID                    int       `gorm:"not null" json:"customer_id"`
	AccountLimit                  float64   `json:"account_limit"`
	PerTransactionLimit           float64   `json:"per_transaction_limit"`
	LastAccountLimit              float64   `json:"last_account_limit"`
	LastPerTransactionLimit       float64   `json:"last_per_transaction_limit"`
	AccountLimitUpdateTime        time.Time `json:"account_limit_update_time"`
	PerTransactionLimitUpdateTime time.Time `json:"per_transaction_limit_update_time"`
	CreatedAt                     time.Time `json:"created_at"`
}

type LimitOffer struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	AccountID           int       `json:"account_id"`
	LimitType           string    `json:"limit_type"`
	NewLimit            float64   `json:"new_limit"`
	OfferActivationTime time.Time `json:"offer_activation_time"`
	OfferExpiryTime     time.Time `json:"offer_expiry_time"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
