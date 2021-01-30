package repository

import (
	"context"
	"github.com/tiwariayush700/auction/models"
)

type BidRepository interface {
	Repository
	FetchBids(ctx context.Context) ([]models.Bid, error)
	GetBidsByAuctionID(ctx context.Context, auctionID uint) ([]models.Bid, error)
}
