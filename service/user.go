package service

import (
	"context"
	"fmt"
	helper "go_catering/helpers"
	"go_catering/resource/model"
	reqModel "go_catering/resource/reqmodel"
	resModel "go_catering/resource/resmodel"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/google/uuid"
	"github.com/joho/godotenv"
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

//CreateSupplier ..
func CreateSupplier(req *reqModel.CreateUserReq) map[string]interface{} {

	newUser := model.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Password: req.Password,
		Address:  req.Address,
		IsSeller: true,
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

	reponse := helper.Message(http.StatusCreated, "Succesfull create supplier")
	reponse["data"] = newUser
	return reponse
}

//Login ...
func Login(req *reqModel.LoginReq) map[string]interface{} {

	filter := bson.M{"$and": []bson.M{
		bson.M{"username": req.Username},
		bson.M{"base.deletedby": ""},
	}}

	filterUser := bson.M{"$and": []bson.M{
		bson.M{"username": req.Username},
		bson.M{"password": req.Password},
		bson.M{"isseller": req.IsSeller},
	}}

	user := model.User{}

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	secretKey := os.Getenv("secret_key")

	err := collectionUser.FindOne(context.TODO(), filter).Decode(&user)
	fmt.Println(err)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Error get users : %v\n", err)
			response := helper.Message(http.StatusNotFound, "User not found")
			response["data"] = nil
			return response
		}
		log.Printf("Error get users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	errFindUser := collectionUser.FindOne(context.TODO(), filterUser).Decode(&user)
	fmt.Println(err)

	if errFindUser != nil {
		if errFindUser == mongo.ErrNoDocuments {
			log.Printf("Error get users : %v\n", errFindUser)
			response := helper.Message(http.StatusNotFound, "Username & Password not Match")
			response["data"] = nil
			return response
		}
		log.Printf("Error get users : %v\n", errFindUser)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	expiredTime := time.Now().Add(1000 * time.Minute)

	claims := &model.Token{
		Username: user.Username,
		ID:       user.ID,
		IsSeller: user.IsSeller,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	fmt.Println(token)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("Error creating jwt users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	response := resModel.LoginRes{
		Username: user.Username,
		Token:    tokenString,
		ID:       user.ID,
		IsSeller: user.IsSeller,
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Login")
	reponse["data"] = response
	return reponse

}
