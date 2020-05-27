package model

import (
	"time"
)

// Transaction ...
type Transaction struct {
	Base
	ID              string            `json:"id"`
	StoreID         string            `json:"store_id"`
	ProductID       string            `json:"product_id"`
	Subsription     bool              `json:"subs"`
	PaymentTime     time.Time         `json:"payment_time"`
	TransactionCode string            `json:"transaction_code"`
	Price           int               `json:"price"`
	Quantity        int               `json:"quantity"`
	Status          StatusTransaction `json:"status"`
	CustomerID      string            `json:"customer_id"`
	IsSubmited      bool              `json:"is_submited"`
}

// StatusTransaction ...
type StatusTransaction int

const (
	//Ordered ...
	Ordered StatusTransaction = iota
	//Success ...
	Success
	//Cancel ...
	Cancel
)
