package repositoryImpl

import (
	"context"
	auctionError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/repository"
	"gorm.io/gorm"
)

type bidRepositoryImpl struct {
	repositoryImpl
}

func (b *bidRepositoryImpl) GetBidsByAuctionID(ctx context.Context, auctionID uint) ([]models.Bid, error) {
	bids := make([]models.Bid, 0)

	err := b.DB.Where("auction_id = ?", auctionID).Find(&bids).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, auctionError.ErrorNotFound
		}
		return nil, err
	}

	return bids, nil
}

func (b *bidRepositoryImpl) FetchBids(ctx context.Context) ([]models.Bid, error) {

	bids := make([]models.Bid, 0)

	err := b.DB.Find(&bids).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, auctionError.ErrorNotFound
		}
		return nil, err
	}

	return bids, nil
}

func NewBidRepositoryImpl(db *gorm.DB) repository.BidRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &bidRepositoryImpl{repoImpl}
}
