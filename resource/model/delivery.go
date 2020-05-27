package model

import (
	"time"
)

//DeliveryProduct ...
type DeliveryProduct struct {
	Base
	ID           string         `json:"id"`
	ProductID    string         `json:"product_id"`
	StoreID      string         `json:"store_id"`
	CustomerID   string         `json:"cust_id"`
	DailyProduct bool           `json:"daily_product"`
	DeliveryDate time.Time      `json:"delivery_date"`
	Status       DeliveryStatus `json:"delivery_status"`
	TodayStatus  DeliveryStatus `json:"today_status"`
}

//DailyDelivery ...
type DailyDelivery struct {
	Base
	ID                string         `json:"id"`
	DeliveryProductID string         `json:"deliv_product_id"`
	DeliveryDate      time.Time      `json:"delivery_date"`
	Status            DeliveryStatus `json:"delivery_status"`
}

// DeliveryStatus ...
type DeliveryStatus int

const (
	//Arrived ...
	Arrived DeliveryStatus = iota
	//Deliver ...
	Deliver
	//Skip ...
	Skip
)
