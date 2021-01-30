package repositoryImpl

import (
	"context"
	auctionError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/repository"
	"gorm.io/gorm"
)

type auctionRepositoryImpl struct {
	repositoryImpl
}

func (a *auctionRepositoryImpl) FetchAuctions(ctx context.Context) ([]models.Auction, error) {
	auctions := make([]models.Auction, 0)

	err := a.DB.Find(&auctions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, auctionError.ErrorNotFound
		}
		return nil, err
	}

	return auctions, nil
}

func (a *auctionRepositoryImpl) GetAuctionsByItemID(ctx context.Context, itemID uint) ([]models.Auction, error) {

	auctions := make([]models.Auction, 0)

	err := a.DB.Where("item_id = ?", itemID).Find(&auctions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, auctionError.ErrorNotFound
		}
		return nil, err
	}

	return auctions, nil
}

func NewAuctionRepositoryImpl(db *gorm.DB) repository.AuctionRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &auctionRepositoryImpl{repoImpl}
}
