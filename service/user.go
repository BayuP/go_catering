package service

import (
	"context"
	helper "go_catering/helpers"
	"go_catering/resource/model"
	reqModel "go_catering/resource/reqmodel"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	uuid "github.com/google/uuid"
)

var collectionUser *mongo.Collection

//UserCollections ...
func UserCollections(m *mongo.Database) {
	collectionUser = m.Collection("user")
}

//CreateUser ..
func CreateUser(req *reqModel.CreateUserReq) map[string]interface{} {

	newUser := model.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Password: req.Password,
		Address:  req.Address,
		IsSeller: false,
		Base: model.Base{
			CreatedTime: time.Now(),
			CreatedBy:   "System",
		},
	}

	_, err := collectionUser.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create user")
	reponse["data"] = newUser
	return reponse
}
