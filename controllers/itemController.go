package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiwariayush700/auction/auth"
	userError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/services"
	"net/http"
)

type itemController struct {
	service     services.ItemService
	app         *app
	authService auth.Service
}

func (controller *itemController) RegisterRoutes() {
	router := controller.app.Router
	itemRouterGroup := router.Group("/items")
	{
		itemRouterGroup.Use(VerifyUserAndServe(controller.authService))
		itemRouterGroup.GET("", controller.GetItems())

		itemRouterGroup.Use(VerifyAdminAndServe(controller.authService))
		itemRouterGroup.POST("", controller.CreateItem())

	}
}

func (controller *itemController) CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		item := &models.Item{}
		err = c.ShouldBind(item)
		if err != nil {
			errRes := userError.NewErrorBadRequest(err, "invalid input")
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		err = controller.service.CreateItem(c, item)
		if err != nil {
			errRes := userError.NewErrorInternal(err, "something went wrong")
			c.JSON(http.StatusInternalServerError, errRes)
			return
		}

		c.JSON(http.StatusOK, item)
	}
}

func (controller *itemController) GetItems() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		status := c.Request.FormValue("status")

		items, err := controller.service.GetItemsForStatus(c, models.ItemStatus(status))
		if err != nil {
			if err == userError.ErrorNotFound {
				errRes := userError.NewErrorNotFound(err, "items not found")
				c.JSON(http.StatusNotFound, errRes)
				return
			}
			errRes := userError.NewErrorInternal(err, "something went wrong")
			c.JSON(http.StatusNotFound, errRes)
			return
		}

		c.JSON(http.StatusOK, items)
	}
}

func NewItemController(service services.ItemService, app *app, authService auth.Service) *itemController {
	return &itemController{service: service, app: app, authService: authService}
}
