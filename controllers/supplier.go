package controller

import (
	helpers "go_catering/helpers"
	req "go_catering/resource/reqmodel"
	service "go_catering/service"

	"github.com/gin-gonic/gin"
)

//CreateSupplier ....
func CreateSupplier(c *gin.Context) {
	var supplier *req.CreateUserReq
	c.BindJSON(&supplier)

	response := service.CreateSupplier(supplier)

	helpers.Respond(c.Writer, response)
}

//CreateStore ....
func CreateStore(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var store *req.ReqCreateStore
	c.BindJSON(&store)

	response := service.CreateStore(idUser, store)

	helpers.Respond(c.Writer, response)
}

//GetAllStore ...
func GetAllStore(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	response := service.GetAllStore(idUser)

	helpers.Respond(c.Writer, response)
}

//GetAllProductByStore ...
func GetAllProductByStore(c *gin.Context) {
	storeID := c.Query("id")
	idUser := c.MustGet("credUser").(string)
	response := service.GetProductByStore(idUser, storeID)

	helpers.Respond(c.Writer, response)
}

//CreateProduct ....
func CreateProduct(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var product *req.ReqCreateProduct
	c.BindJSON(&product)

	response := service.CreateProduct(idUser, product)

	helpers.Respond(c.Writer, response)
}
