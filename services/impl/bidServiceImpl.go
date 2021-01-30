package serviceImpl

import (
	"context"
	auctionError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/repository"
	"github.com/tiwariayush700/auction/services"
	"gorm.io/gorm"
)

type bidServiceImpl struct {
	repository repository.BidRepository
}

func (b *bidServiceImpl) GetBidsByAuctionID(ctx context.Context, auctionID uint) ([]models.Bid, error) {
	return b.repository.GetBidsByAuctionID(ctx, auctionID)
}

func (b *bidServiceImpl) FetchBids(ctx context.Context) ([]models.Bid, error) {
	return b.repository.FetchBids(ctx)
}

func (b *bidServiceImpl) GetBidByID(ctx context.Context, bidID uint) (*models.Bid, error) {
	bid := &models.Bid{}
	err := b.repository.Get(ctx, bid, bidID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, auctionError.ErrorNotFound
		}
		return nil, err
	}

	return bid, err
}

func (b *bidServiceImpl) CreateBid(ctx context.Context, bid *models.Bid) error {
	return b.repository.Create(ctx, bid)
}

func NewBidServiceImpl(repository repository.BidRepository) services.BidService {
	return &bidServiceImpl{repository: repository}
}
