package http

import (
	_ "bank-api/docs"
	"bank-api/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(service *service.Service) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	controller := newController(service)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		client := api.Group("/client")
		{
			client.POST("/create", controller.createClient)
			client.GET("/list", controller.getClients)
			client.GET("/", controller.getClient)
		}

		account := api.Group("/account")
		{
			account.POST("/create", controller.createAccount)
			account.GET("/list", controller.getAccounts)
			account.GET("/", controller.getAccount)
		}

		transaction := api.Group("/transaction")
		{
			transaction.POST("/create", controller.createTransaction)
			transaction.GET("/list", controller.getTransactions)
			transaction.GET("/", controller.getTransaction)
		}
	}
	return router
}
