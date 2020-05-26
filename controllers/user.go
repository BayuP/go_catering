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
