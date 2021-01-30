package serviceImpl

import (
	"github.com/tiwariayush700/auction/repository"
	"github.com/tiwariayush700/auction/services"
)

type auctionServiceImpl struct {
	repository repository.AuctionRepository
}

func NewAuctionServiceImpl(repository repository.AuctionRepository) services.AuctionService {
	return &auctionServiceImpl{repository: repository}
}
