package model

import (
	"time"
)

// Transaction ...
type Transaction struct {
	Base
	ID          string            `json:"id"`
	StoreID     string            `json:"store_id"`
	ProductID   string            `'json:"product_id"`
	Subsription bool              `json:"subs"`
	PaymentTime time.Time         `json:"payment_time"`
	Price       int               `json:"price"`
	Status      StatusTransaction `json:"status"`
}

// StatusTransaction ...
type StatusTransaction int

const (
	//Pending ...
	Pending StatusTransaction = iota
	//Success ...
	Success
	//Cancel ...
	Cancel
)
