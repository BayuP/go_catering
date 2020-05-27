package routers

import (
	controller "go_catering/controllers"
	"go_catering/middleware"
	db "go_catering/resource"

	"github.com/gin-gonic/gin"
)

//SetupRouter ...
func SetupRouter() *gin.Engine {

	r := gin.Default()

	//connecting to db
	db.Connect()
	// Routing endpoint

	user := r.Group("/api/user")
	supplier := r.Group("/api/supplier")
	// category := r.Group("/api/category")

	user.POST("/login", controller.LoginUser)
	user.POST("/", controller.CreateUser)
	supplier.POST("/", controller.CreateSupplier)

	supplier.Use(middleware.AuthMiddlewares())
	{
		supplier.POST("/store", controller.CreateStore)
		supplier.GET("/store/all", controller.GetAllStore)
		supplier.GET("/store/product", controller.GetAllProductByStore)
		supplier.POST("/product", controller.CreateProduct)
		supplier.GET("/delivery/all", controller.GetAllDeliv)
		supplier.PUT("/delivery", controller.UpdateDelivery)
	}
	user.Use(middleware.AuthMiddlewaresUser())
	{
		user.GET("/product", controller.AllProduct)
		user.POST("/trx", controller.CreateTransaction)
		user.PUT("/trx", controller.UpdateTransaction)
	}

	return r
}
