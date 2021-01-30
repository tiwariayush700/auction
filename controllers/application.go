package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	authImpl "github.com/tiwariayush700/auction/auth/impl"
	"github.com/tiwariayush700/auction/config"
	"github.com/tiwariayush700/auction/models"
	"gorm.io/gorm"
)

// App structure for tenant microservice
type app struct {
	Config *config.Config
	DB     *gorm.DB //set from main.go
	Router *gin.Engine
}

func NewApp(config *config.Config, db *gorm.DB, router *gin.Engine) *app {
	return &app{
		Config: config,
		DB:     db,
		Router: router,
	}
}

func (app *app) Start() {

	app.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "HEAD", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	//repositories

	//services
	_ = authImpl.NewAuthService(app.Config.AuthSecret)

	//controllers

	//register routes

	app.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	logrus.Info("Application loaded successfully ")
	logrus.Fatal(app.Router.Run(":" + app.Config.Port))

}

func (app *app) Migrate() error {
	if err := app.DB.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	if err := app.DB.AutoMigrate(&models.Item{}); err != nil {
		return err
	}

	if err := app.DB.AutoMigrate(&models.Auction{}); err != nil {
		return err
	}

	if err := app.DB.AutoMigrate(&models.Bid{}); err != nil {
		return err
	}

	return nil
}
