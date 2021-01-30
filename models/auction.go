package models

import (
	"gorm.io/gorm"
	"time"
)

type Auction struct {
	gorm.Model
	StartPrice float64   `json:"start_price" gorm:"default:0;not null"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`

	//foreign keys
	ItemID uint `json:"item_id"`

	Item *Item `json:"-"`
}
