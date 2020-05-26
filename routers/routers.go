package routers

import (
	controller "go_catering/controllers"
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
	// product := r.Group("/api/product")
	// category := r.Group("/api/category")

	// user.POST("/login", controller.LoginUser)
	user.POST("/", controller.CreateUser)
	// user.Use(middleware.AuthMiddlewares())
	// {

	// 	user.GET("/all", v1Controller.GetAllUsers)
	// 	//v1.GET("/user", v1Controller.GetUser)
	// 	user.POST("/", v1Controller.CreateUser)
	// 	user.PUT("/", v1Controller.UpdateUser)
	// 	user.DELETE("/", v1Controller.DeleteUser)
	// 	user.GET("/", v1Controller.GetUser)
	// 	user.POST("/changepass", v1Controller.ChangePassword)
	// }

	// product.Use(middleware.AuthMiddlewares())
	// {
	// 	product.POST("/create_product", v1Controller.CreateProduct)
	// 	product.GET("/all", v1Controller.GetAllProduct)
	// 	product.GET("/", v1Controller.GetProduct)
	// 	product.PUT("/", v1Controller.UpdateProductByID)
	// 	product.DELETE("/", v1Controller.SoftDeleteByID)
	// }

	// category.Use(middleware.AuthMiddlewares())
	// {
	// 	category.POST("/create_category", v1Controller.CreateCategory)
	// 	category.GET("/all", v1Controller.GetAllCategory)
	// 	category.GET("/", v1Controller.GetCategory)
	// 	category.PUT("/", v1Controller.UpdateCategoryByID)
	// 	category.DELETE("/", v1Controller.DeleteCatByID)
	// }

	return r
}