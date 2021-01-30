package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiwariayush700/auction/auth"
	"github.com/tiwariayush700/auction/services"
)

type auctionController struct {
	service     services.AuctionService
	app         *app
	authService auth.Service
}

func (controller *auctionController) RegisterRoutes() {
	router := controller.app.Router
	itemRouterGroup := router.Group("/auctions")
	{
		itemRouterGroup.Use(VerifyUserAndServe(controller.authService))
		itemRouterGroup.GET("", controller.GetAuctions())

		itemRouterGroup.Use(VerifyAdminAndServe(controller.authService))
		itemRouterGroup.POST("", controller.CreateAuction())

	}
}

func (controller *auctionController) GetAuctions() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (controller *auctionController) CreateAuction() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func NewAuctionController(service services.AuctionService, app *app, authService auth.Service) *auctionController {
	return &auctionController{service: service, app: app, authService: authService}
}
