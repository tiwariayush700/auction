package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tiwariayush700/auction/auth"
	auctionError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/services"
	"net/http"
	"strconv"
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
	bidRouterGroup := router.Group("/bids")
	{
		bidRouterGroup.Use(VerifyUserAndServe(controller.authService))
		bidRouterGroup.POST("", controller.CreateBid())
		bidRouterGroup.GET("/:bid_id", controller.GetBidResult())

		bidRouterGroup.Use(VerifyAdminAndServe(controller.authService))
		bidRouterGroup.GET("", controller.GetBids())

	}
}

func (controller *bidController) GetBidResult() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		bidIdParam, ok := c.Params.Get("bid_id")
		if !ok {
			c.JSON(http.StatusBadRequest, auctionError.NewErrorBadRequest(err, "Invalid Bid ID"))
			return
		}

		bidId, err := strconv.ParseUint(bidIdParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, auctionError.NewErrorBadRequest(err, "err parsing Bid ID"))
			return
		}

		bid, err := controller.service.GetBidByID(c, uint(bidId))
		if err != nil {
			if err == auctionError.ErrorNotFound {
				c.JSON(http.StatusNotFound, auctionError.NewErrorNotFound(err, "No bid found for the given bid id"))
				return
			}
			c.JSON(http.StatusInternalServerError, auctionError.NewErrorInternal(err, "something went wrong"))
			return
		}

		auction, err := controller.auctionService.GetAuctionByID(c, bid.AuctionID)
		if err != nil {
			if err == auctionError.ErrorNotFound {
				c.JSON(http.StatusNotFound, auctionError.NewErrorNotFound(err, "No auction found for the given auction id"))
				return
			}
			c.JSON(http.StatusInternalServerError, auctionError.NewErrorInternal(err, "something went wrong"))
			return
		}

		bids, err := controller.service.GetBidsByAuctionID(c, bid.AuctionID)
		if err != nil {
			if err == auctionError.ErrorNotFound {
				c.JSON(http.StatusNotFound, auctionError.NewErrorNotFound(err, "No bid found for the given auction id"))
				return
			}
			c.JSON(http.StatusInternalServerError, auctionError.NewErrorInternal(err, "something went wrong"))
			return
		}

		maximumBid := float64(0)
		userID := uint(0)
		for _, val := range bids {
			if val.Amount > maximumBid {
				maximumBid = val.Amount
				userID = val.UserID
			}
		}

		timeNow := time.Now()
		if timeNow.Before(auction.StartTime) {
			c.JSON(http.StatusBadRequest, auctionError.NewErrorBadRequest(auctionError.ErrorBadRequest,
				fmt.Sprintf("Auction has not started yet. It will start at %v", auction.StartTime)))
			return
		}

		if timeNow.After(auction.EndTime) {
			//TODO update item status to unbiddable or sold
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Auction is complete. Max Bid amount was %0.2f against user with userID : %v", maximumBid, userID),
				"data":    bid,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Auction is ongoing. Current max bid is %0.2f by user with userID : %v", maximumBid, userID),
			"data":    bid,
		})
	}
}

func (controller *bidController) GetBids() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		bids, err := controller.service.FetchBids(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, auctionError.NewErrorInternal(err, "something went wrong"))
			return
		}

		c.JSON(http.StatusOK, bids)
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
