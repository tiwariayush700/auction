package repositoryImpl

import (
	"github.com/tiwariayush700/auction/repository"
	"gorm.io/gorm"
)

type auctionRepositoryImpl struct {
	repositoryImpl
}

func NewAuctionRepositoryImpl(db *gorm.DB) repository.AuctionRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &auctionRepositoryImpl{repoImpl}
}
