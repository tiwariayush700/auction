package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name   string     `json:"name" gorm:"type:text;not null"`
	Status ItemStatus `json:"status" binding:"required,oneof=biddable unBiddable" gorm:"type:text;check:status = 'biddable' or status = 'unBiddable';not null"`
}

type ItemStatus string

const (
	ItemStatusBiddable   = ItemStatus("biddable")
	ItemStatusUnBiddable = ItemStatus("unBiddable")
)
