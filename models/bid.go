package models

import "gorm.io/gorm"

type Bid struct {
	gorm.Model
	Amount float64 `json:"amount" gorm:"default:0;not null"`

	//foreign keys
	UserID    uint `json:"user_id"`
	AuctionID uint `json:"auction_id"`

	User    *User    `json:"-"`
	Auction *Auction `json:"-"`
}

type BidRequest struct {
	AuctionID uint    `json:"auction_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}
