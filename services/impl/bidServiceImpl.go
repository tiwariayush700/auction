package serviceImpl

import (
	"context"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/repository"
	"github.com/tiwariayush700/auction/services"
)

type bidServiceImpl struct {
	repository repository.BidRepository
}

func (b *bidServiceImpl) CreateBid(ctx context.Context, bid *models.Bid) error {
	return b.repository.Create(ctx, bid)
}

func NewBidServiceImpl(repository repository.BidRepository) services.BidService {
	return &bidServiceImpl{repository: repository}
}
