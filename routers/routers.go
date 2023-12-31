package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.Home)
	router.POST("/order", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.GET("/order/:id", controllers.GetOrderById)
	router.PUT("/order/:id", controllers.UpdateOrder)
	router.DELETE("/order/:id", controllers.DeleteOrder)

	return router
}
