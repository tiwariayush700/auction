package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiwariayush700/auction/auth"
	"github.com/tiwariayush700/auction/services"
)

type bidController struct {
	service        services.BidService
	auctionService services.AuctionService
	app            *app
	authService    auth.Service
}

func (controller *bidController) RegisterRoutes() {
	router := controller.app.Router
	itemRouterGroup := router.Group("/bids")
	{
		itemRouterGroup.Use(VerifyUserAndServe(controller.authService))
		itemRouterGroup.GET("", controller.GetBids())

		itemRouterGroup.Use(VerifyAdminAndServe(controller.authService))
		itemRouterGroup.POST("", controller.CreateBid())

	}
}

func (controller *bidController) GetBids() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (controller *bidController) CreateBid() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func NewBidController(bidService services.BidService, auctionService services.AuctionService, app *app, authService auth.Service) *bidController {
	return &bidController{service: bidService, auctionService: auctionService, app: app, authService: authService}
}
