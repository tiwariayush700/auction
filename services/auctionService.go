package services

import (
	"context"
	"github.com/tiwariayush700/auction/models"
)

type AuctionService interface {
	CreateAuction(ctx context.Context, auction *models.Auction) error
	GetAuctionsByItemID(ctx context.Context, itemID uint) ([]models.Auction, error)
}
