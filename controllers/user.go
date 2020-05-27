package controller

import (
	helpers "go_catering/helpers"
	req "go_catering/resource/reqmodel"
	service "go_catering/service"

	"github.com/gin-gonic/gin"
)

//CreateUser ....
func CreateUser(c *gin.Context) {
	var user *req.CreateUserReq
	c.BindJSON(&user)

	response := service.CreateUser(user)

	helpers.Respond(c.Writer, response)
}

//LoginUser ..
func LoginUser(c *gin.Context) {
	var login *req.LoginReq
	c.BindJSON(&login)

	response := service.Login(login)

	helpers.Respond(c.Writer, response)
}

//AllProduct ..
func AllProduct(c *gin.Context) {

	response := service.GetAllProduct()

	helpers.Respond(c.Writer, response)
}

//CreateTransaction ....
func CreateTransaction(c *gin.Context) {
	idUser := c.MustGet("credUser").(string)
	var trx *req.TransactionReq
	c.BindJSON(&trx)

	response := service.Transaction(idUser, trx)

	helpers.Respond(c.Writer, response)
}

//UpdateTransaction ....
func UpdateTransaction(c *gin.Context) {
	var trx *req.UpdateTrxReq
	c.BindJSON(&trx)

	response := service.UpdateTransaction(trx)

	helpers.Respond(c.Writer, response)
}
