package service

import (
	"context"
	helper "go_catering/helpers"
	"go_catering/resource/model"
	reqModel "go_catering/resource/reqmodel"
	resModel "go_catering/resource/resmodel"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	uuid "github.com/google/uuid"
)

var collectionStore *mongo.Collection
var collectionProduct *mongo.Collection

//StoreCollections ...
func StoreCollections(m *mongo.Database) {
	collectionStore = m.Collection("store")
}

//ProductCollections ...
func ProductCollections(m *mongo.Database) {
	collectionProduct = m.Collection("product")
}

//CreateStore ..
func CreateStore(userID string, req *reqModel.ReqCreateStore) map[string]interface{} {

	newStore := model.Store{
		ID:           uuid.New().String(),
		UserID:       userID,
		StoreName:    req.StoreName,
		StoreAddress: req.StoreAddress,
		SellingArea:  req.SellingArea,
		Base: model.Base{
			CreatedTime: time.Now(),
			CreatedBy:   userID,
		},
	}

	_, err := collectionStore.InsertOne(context.TODO(), newStore)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create store")
	reponse["data"] = newStore
	return reponse
}

//GetAllStore ...
func GetAllStore(userID string) map[string]interface{} {
	filter := bson.M{"$and": []bson.M{
		bson.M{"userid": userID},
		bson.M{"base.deletedby": ""},
	}}

	result := []resModel.StoreRes{}
	cursor, err := collectionStore.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error when getting all store %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	for cursor.Next(context.TODO()) {
		var store *model.Store
		cursor.Decode(&store)
		storeRes := resModel.StoreRes{
			ID:           store.ID,
			StoreName:    store.StoreName,
			StoreAddress: store.StoreAddress,
			SellingArea:  store.SellingArea,
		}
		result = append(result, storeRes)
	}

	reponse := helper.Message(http.StatusOK, "Succesfull get All Store")
	reponse["data"] = result
	return reponse
}

//CreateProduct ..
func CreateProduct(userID string, req *reqModel.ReqCreateProduct) map[string]interface{} {

	newProduct := model.Product{
		ID:           uuid.New().String(),
		Name:         req.Name,
		StoreID:      req.StoreID,
		Price:        req.Price,
		Stock:        req.Stock,
		DailyProduct: req.DailyProduct,
		Base: model.Base{
			CreatedTime: time.Now(),
			CreatedBy:   userID,
		},
	}

	_, err := collectionProduct.InsertOne(context.TODO(), newProduct)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusCreated, "Succesfull create product")
	reponse["data"] = newProduct
	return reponse
}

//GetProductByStore ..
func GetProductByStore(userID string, storeID string) map[string]interface{} {

	filter := bson.M{"$and": []bson.M{
		bson.M{"base.createdby": userID},
		bson.M{"storeid": storeID},
		bson.M{"base.deletedby": ""},
	}}

	result := []resModel.ProductByStoreRes{}

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
		productRes := resModel.ProductByStoreRes{
			ID:           product.ID,
			Name:         product.Name,
			Price:        product.Price,
			Stock:        product.Stock,
			DailyProduct: product.DailyProduct,
		}
		result = append(result, productRes)
	}

	reponse := helper.Message(http.StatusOK, "Succesfull get All Product")
	reponse["data"] = result
	return reponse
}
