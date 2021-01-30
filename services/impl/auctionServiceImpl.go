package serviceImpl

import (
	"context"
	auctionError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/repository"
	"github.com/tiwariayush700/auction/services"
	"gorm.io/gorm"
)

type auctionServiceImpl struct {
	repository repository.AuctionRepository
}

func (a *auctionServiceImpl) UpdateAuctionPrice(ctx context.Context, amount float64, auctionID uint) error {
	return a.repository.UpdateAuctionPrice(ctx, amount, auctionID)
}

func (a *auctionServiceImpl) GetAuctionsByItemID(ctx context.Context, itemID uint) ([]models.Auction, error) {
	if itemID > 0 {
		return a.repository.GetAuctionsByItemID(ctx, itemID)
	}

	return a.repository.FetchAuctions(ctx)
}

func (a *auctionServiceImpl) GetAuctionByID(ctx context.Context, auctionID uint) (*models.Auction, error) {

	auction := &models.Auction{}
	err := a.repository.Get(ctx, auction, auctionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, auctionError.ErrorNotFound
		}
		return nil, err
	}

	return auction, err
}

func (a *auctionServiceImpl) CreateAuction(ctx context.Context, auction *models.Auction) error {
	return a.repository.Create(ctx, auction)
}

func NewAuctionServiceImpl(repository repository.AuctionRepository) services.AuctionService {
	return &auctionServiceImpl{repository: repository}
}
