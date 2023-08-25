package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/beparwaah/assignmentVegapay/models"
	"github.com/beparwaah/assignmentVegapay/services"
	"github.com/gin-gonic/gin"
)

func CreateLimitOffer(c *gin.Context) {
	var offer models.LimitOffer
	if err := c.ShouldBindJSON(&offer); err != nil {
		// Responding with Invalid request
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	offer.OfferActivationTime = time.Now()
	offer.Status = "PENDING"

	err := services.CreateLimitOffer(&offer)
	if err != nil {
		// error while creating error
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating limit offer"})
		return
	}

	// request successful
	c.JSON(http.StatusCreated, gin.H{"message": "Limit offer created successfully"})
}
func ListActiveLimitOffers(c *gin.Context) {
	accountIDStr := c.Query("account_id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		// Invalid account ID
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	activeDateStr := c.Query("active_date")
	activeDate, err := time.Parse("2006-01-02", activeDateStr)
	if err != nil {
		//passing on current time
		activeDate = time.Now()
	}

	offers, err := services.ListActiveLimitOffers(accountID, activeDate)
	if err != nil {
		//Error fetching active offers
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching active offers"})
		return
	}

	//successfull request
	c.JSON(http.StatusOK, offers)
}
func UpdateLimitOfferStatus(c *gin.Context) {

	offerID := c.Param("id")
	fmt.Println("ofer id is offerID", offerID)
	offerIDUint, err := strconv.ParseInt(offerID, 10, 64)
	if err != nil {
		// Invalid offer ID
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offer ID"})
		return
	}

	status := c.Param("status")

	offer, err := services.GetLimitOfferByID(uint(offerIDUint))
	if err != nil {
		// Error fetching offer
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching offer"})
		return
	}

	if offer == nil {
		//Offer not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	err = services.UpdateLimitOfferStatus(offer, status)
	if err != nil {
		// Error updating offer status
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating offer status"})
		return
	}

	// successful request
	c.JSON(http.StatusOK, gin.H{"message": "Offer status updated successfully"})
}
func CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		// Invalid request payload
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := services.CreateAccount(&account)
	if err != nil {
		// Error creating account
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating account"})
		return
	}

	// request successful
	c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully"})
}
func GetAccountByID(c *gin.Context) {
	accountIDStr := c.Query("account_id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		//Invalid account ID
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	account, err := services.GetAccountByID(accountID)
	if err != nil {
		// Error fetching account
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching account"})
		return
	}

	if account == nil {
		// Account not found
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	// request successful
	c.JSON(http.StatusOK, account)
}
