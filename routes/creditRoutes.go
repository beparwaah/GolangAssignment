package routes

import (
	"github.com/beparwaah/assignmentVegapay/controller"
	"github.com/gin-gonic/gin"
)

func CreditRoutes(router *gin.Engine) {
	//create limit offer
	router.POST("/limit-offer/create", controller.CreateLimitOffer)
	//list of active limit offers
	router.GET("/limit-offer/list-active", controller.ListActiveLimitOffers)
	//update limit offer status
	router.PUT("/limit-offer/update-status/:id/:status", controller.UpdateLimitOfferStatus)
	//create account
	router.POST("/account/create", controller.CreateAccount)
	//get account
	router.GET("/account/get", controller.GetAccountByID)

}
