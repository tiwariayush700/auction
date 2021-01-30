package services

import (
	"context"
	"github.com/tiwariayush700/auction/models"
)

type BidService interface {
	CreateBid(ctx context.Context, bid *models.Bid) error
}
