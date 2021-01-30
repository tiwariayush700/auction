package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiwariayush700/auction/auth"
	auctionError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/services"
	"net/http"
	"time"
)

type auctionController struct {
	service     services.AuctionService
	itemService services.ItemService
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

		_, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		params := &models.AuctionRequest{}
		err = c.ShouldBind(&params)
		if err != nil {
			errRes := auctionError.NewErrorBadRequest(err, "invalid input")
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		startTime, err := time.Parse(time.RFC3339, params.StartTime)
		if err != nil {
			errRes := auctionError.NewErrorBadRequest(err, "invalid start time")
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		endTime, err := time.Parse(time.RFC3339, params.EndTime)
		if err != nil {
			errRes := auctionError.NewErrorBadRequest(err, "invalid end time")
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		auction := &models.Auction{
			StartPrice: params.StartPrice,
			ItemID:     params.ItemID,
			StartTime:  startTime,
			EndTime:    endTime,
		}

		item, err := controller.itemService.GetItemByID(c, auction.ItemID)
		if err != nil {
			if err == auctionError.ErrorNotFound {
				c.JSON(http.StatusNotFound, auctionError.NewErrorNotFound(err, "Item for the given id not found"))
				return
			}
			c.JSON(http.StatusInternalServerError, auctionError.NewErrorInternal(err, "something went wrong"))
			return
		}

		if item.Status != models.ItemStatusBiddable {
			c.JSON(http.StatusBadRequest, auctionError.NewErrorBadRequest(err, "cannot auction unBiddable item"))
			return
		}

		err = controller.service.CreateAuction(c, auction)
		if err != nil {
			errRes := auctionError.NewErrorInternal(err, "something went wrong")
			c.JSON(http.StatusInternalServerError, errRes)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Auction created successfully",
			"auction": auction,
		})
	}
}

func NewAuctionController(service services.AuctionService, itemService services.ItemService, app *app, authService auth.Service) *auctionController {
	return &auctionController{service: service, itemService: itemService, app: app, authService: authService}
}
