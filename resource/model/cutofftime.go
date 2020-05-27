package model

import (
	"time"
)

//CutOffTime ...
type CutOffTime struct {
	Base
	ID            string    `json:"id"`
	LastOrderTime time.Time `json:"last_order_time"`
}
