package controller

import (
	"github.com/tiwariayush700/auction/auth"
	"github.com/tiwariayush700/auction/services"
)

type itemController struct {
	service     services.ItemService
	app         *app
	authService auth.Service
}

func (u *itemController) RegisterRoutes() {
	//router := u.app.Router
	//itemRouterGroup := router.Group("/items")
	//{
	//
	//}
}

func NewItemController(service services.ItemService, app *app, authService auth.Service) *itemController {
	return &itemController{service: service, app: app, authService: authService}
}
