package services

import (
	"context"
	"github.com/tiwariayush700/auction/models"
)

type AuctionService interface {
	CreateAuction(ctx context.Context, auction *models.Auction) error
	GetAuctionsByItemID(ctx context.Context, itemID uint) ([]models.Auction, error)
	GetAuctionByID(ctx context.Context, auctionID uint) (*models.Auction, error)
	UpdateAuctionPrice(ctx context.Context, amount float64, auctionID uint) error
}
