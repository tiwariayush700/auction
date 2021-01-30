package repository

import (
	"context"
	"github.com/tiwariayush700/auction/models"
)

type AuctionRepository interface {
	Repository
	GetAuctionsByItemID(ctx context.Context, itemID uint) ([]models.Auction, error)
	FetchAuctions(ctx context.Context) ([]models.Auction, error)
	UpdateAuctionPrice(ctx context.Context, amount float64, auctionID uint) error
}

