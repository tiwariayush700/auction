package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tiwariayush700/auction/auth"
	auctionError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/services"
	"net/http"
	"time"
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

		//itemRouterGroup.Use(VerifyAdminAndServe(controller.authService))
		itemRouterGroup.POST("", controller.CreateBid())

	}
}

func (controller *bidController) GetBids() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (controller *bidController) CreateBid() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		params := &models.BidRequest{}
		err = c.ShouldBind(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, auctionError.NewErrorBadRequest(err, "invalid input"))
			return
		}

		auction, err := controller.auctionService.GetAuctionByID(c, params.AuctionID)
		if err != nil {
			if err == auctionError.ErrorNotFound {
				c.JSON(http.StatusNotFound, auctionError.NewErrorNotFound(err, "auction for which you applied does not exist"))
				return
			}
		}

		err = validateBid(params, auction)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		bid := &models.Bid{
			Amount:    params.Amount,
			UserID:    userId,
			AuctionID: params.AuctionID,
		}

		err = controller.service.CreateBid(c, bid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, auctionError.NewErrorInternal(err, "something went wrong"))
			return
		}

		err = controller.auctionService.UpdateAuctionPrice(c, bid.Amount, bid.AuctionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, auctionError.NewErrorInternal(err, "err updating auction price"))
			return
		}

		c.JSON(http.StatusOK, bid)
	}
}

func validateBid(params *models.BidRequest, auction *models.Auction) error {

	if params.Amount <= auction.StartPrice {
		return auctionError.NewErrorBadRequest(auctionError.ErrorBadRequest,
			fmt.Sprintf("You have to bid more than %0.2f", auction.StartPrice))
	}

	timeNow := time.Now()

	if timeNow.Before(auction.StartTime) {
		return auctionError.NewErrorBadRequest(auctionError.ErrorBadRequest,
			fmt.Sprintf("Auction has not started yet. It will start at %v", auction.StartTime))
	}

	if timeNow.After(auction.EndTime) {
		return auctionError.NewErrorBadRequest(auctionError.ErrorBadRequest,
			"The auction is over")
	}

	return nil
}

func NewBidController(bidService services.BidService, auctionService services.AuctionService, app *app, authService auth.Service) *bidController {
	return &bidController{service: bidService, auctionService: auctionService, app: app, authService: authService}
}
