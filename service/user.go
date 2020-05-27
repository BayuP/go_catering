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
var collectionTransaction *mongo.Collection
var collectionDelivery *mongo.Collection

//UserCollections ...
func UserCollections(m *mongo.Database) {
	collectionUser = m.Collection("user")
}

//TransactionCollections ...
func TransactionCollections(m *mongo.Database) {
	collectionTransaction = m.Collection("transaction")
}

//DeliveryCollections ...
func DeliveryCollections(m *mongo.Database) {
	collectionDelivery = m.Collection("deliveryProduct")
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
		log.Printf("Error when inserting new transaction : %v\n", err)
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

//GetAllProduct ..
func GetAllProduct() map[string]interface{} {
	filter := bson.M{"base.deletedby": ""}

	result := []resModel.AllProduct{}

	cursor, err := collectionProduct.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error when getting all product %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	for cursor.Next(context.TODO()) {
		var product *model.Product
		cursor.Decode(&product)
		store := model.Store{}
		filterStore := bson.M{"$and": []bson.M{
			bson.M{"id": product.StoreID},
			bson.M{"base.deletedby": ""},
		}}
		err := collectionStore.FindOne(context.TODO(), filterStore).Decode(&store)

		if err != nil {
			log.Printf("Error when getting all product %v\n", err)
			response := helper.Message(http.StatusInternalServerError, "Someting wrong")
			response["data"] = nil
			return response
		}

		productRes := resModel.AllProduct{
			ID:           product.ID,
			Name:         product.Name,
			Price:        product.Price,
			Stock:        product.Stock,
			DailyProduct: product.DailyProduct,
			StoreName:    store.StoreName,
			StoreID:      product.StoreID,
		}
		result = append(result, productRes)
	}

	reponse := helper.Message(http.StatusOK, "Succesfull get All Product")
	reponse["data"] = result
	return reponse
}

//Transaction ..
func Transaction(userID string, req *reqModel.TransactionReq) map[string]interface{} {

	currentTime := time.Now()
	code := currentTime.Format("20060102 150405 ")
	guidTrx := uuid.New().String()
	newTransaction := model.Transaction{
		ID:              guidTrx,
		StoreID:         req.StoreID,
		ProductID:       req.ProductID,
		Subsription:     req.Subsription,
		Price:           req.Price,
		Quantity:        req.Quantity,
		CustomerID:      userID,
		TransactionCode: code + guidTrx[0:6],
		Status:          1,
		IsSubmited:      false,
		Base: model.Base{
			CreatedTime: time.Now(),
			CreatedBy:   userID,
		},
	}

	_, err := collectionTransaction.InsertOne(context.TODO(), newTransaction)

	if err != nil {
		fmt.Print(err)
		log.Printf("Error when inserting new transcation : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create transaction")
	reponse["data"] = newTransaction
	return reponse
}

//UpdateTransaction ..
func UpdateTransaction(req *reqModel.UpdateTrxReq) map[string]interface{} {
	const CutOffTime = 22

	filter := bson.M{"$and": []bson.M{
		bson.M{"transactioncode": req.TransactionCode},
		bson.M{"base.deletedby": ""},
	}}

	if req.StatusTrx == 2 {
		newData := bson.M{
			"$set": bson.M{
				"paymenttime":      time.Now(),
				"status":           req.StatusTrx,
				"issubmited":       true,
				"base.updatedtime": time.Now(),
				"base.updatedby":   "System",
			},
		}

		result, err := collectionTransaction.UpdateOne(context.TODO(), filter, newData)

		if err != nil {
			log.Printf("Error when updating product : %v\n", err)
			response := helper.Message(http.StatusInternalServerError, "Someting wrong")
			response["data"] = nil
			return response
		}

		if result.MatchedCount == 0 {
			response := helper.Message(http.StatusNotFound, "Not found Document")
			response["data"] = nil
			return response
		}

		trx := model.Transaction{}

		errTrx := collectionTransaction.FindOne(context.TODO(), filter).Decode(&trx)

		if errTrx != nil {
			if err == mongo.ErrNoDocuments {
				response := helper.Message(http.StatusNotFound, "Not found document")
				response["data"] = nil
				return response
			}
			log.Printf("Error when get Product : %v\n", err)
			response := helper.Message(http.StatusInternalServerError, "Someting wrong")
			response["data"] = nil
			return response
		}
		if time.Now().Hour() > CutOffTime {
			newDelivery := model.DeliveryProduct{
				ID:           uuid.New().String(),
				StoreID:      trx.StoreID,
				ProductID:    trx.ProductID,
				CustomerID:   trx.CustomerID,
				DailyProduct: trx.Subsription,
				DeliveryDate: time.Now().AddDate(0, 0, 1),
				Base: model.Base{
					CreatedTime: time.Now(),
					CreatedBy:   "System",
				},
			}

			_, err := collectionDelivery.InsertOne(context.TODO(), newDelivery)

			if err != nil {
				log.Printf("Error when inserting new Delivery : %v\n", err)
				response := helper.Message(http.StatusInternalServerError, "Someting wrong")
				response["data"] = nil
				return response
			}

		} else {
			fmt.Printf("masok")
			newDelivery := model.DeliveryProduct{
				ID:           uuid.New().String(),
				StoreID:      trx.StoreID,
				ProductID:    trx.ProductID,
				CustomerID:   trx.CustomerID,
				DailyProduct: trx.Subsription,
				DeliveryDate: time.Now(),
				Base: model.Base{
					CreatedTime: time.Now(),
					CreatedBy:   "System",
				},
			}

			_, err := collectionDelivery.InsertOne(context.TODO(), newDelivery)

			if err != nil {
				log.Printf("Error when inserting new Delivery : %v\n", err)
				response := helper.Message(http.StatusInternalServerError, "Someting wrong")
				response["data"] = nil
				return response
			}
		}

	}

	if req.StatusTrx == 3 {
		newData := bson.M{
			"$set": bson.M{
				"status":           req.StatusTrx,
				"issubmited":       true,
				"base.updatedtime": time.Now(),
				"base.updatedby":   "System",
			},
		}

		result, err := collectionTransaction.UpdateOne(context.TODO(), filter, newData)

		fmt.Println(result)
		if err != nil {
			log.Printf("Error when updating product : %v\n", err)
			response := helper.Message(http.StatusInternalServerError, "Someting wrong")
			response["data"] = nil
			return response
		}

		if result.MatchedCount == 0 {
			response := helper.Message(http.StatusNotFound, "Not found Document")
			response["data"] = nil
			return response
		}

	}

	reponse := helper.Message(http.StatusOK, "Succesfull Update Transaction")
	reponse["data"] = nil
	return reponse

}
