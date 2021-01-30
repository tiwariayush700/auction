package serviceImpl

import (
	"context"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/repository"
	"github.com/tiwariayush700/auction/services"
)

type auctionServiceImpl struct {
	repository repository.AuctionRepository
}

func (a *auctionServiceImpl) GetAuctionsByItemID(ctx context.Context, itemID uint) ([]models.Auction, error) {
	if itemID > 0 {
		return a.repository.GetAuctionsByItemID(ctx, itemID)
	}

	return a.repository.FetchAuctions(ctx)
}

func (a *auctionServiceImpl) CreateAuction(ctx context.Context, auction *models.Auction) error {
	return a.repository.Create(ctx, auction)
}

func NewAuctionServiceImpl(repository repository.AuctionRepository) services.AuctionService {
	return &auctionServiceImpl{repository: repository}
}
