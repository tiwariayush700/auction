package services

import (
	"context"
	"github.com/tiwariayush700/auction/models"
)

type BidService interface {
	CreateBid(ctx context.Context, bid *models.Bid) error
	GetBidByID(ctx context.Context, bidID uint) (*models.Bid, error)
	FetchBids(ctx context.Context) ([]models.Bid, error)
	GetBidsByAuctionID(ctx context.Context, auctionID uint) ([]models.Bid, error)
}
