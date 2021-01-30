package repositoryImpl

import (
	"github.com/tiwariayush700/auction/repository"
	"gorm.io/gorm"
)

type bidRepositoryImpl struct {
	repositoryImpl
}

func NewBidRepositoryImpl(db *gorm.DB) repository.BidRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &bidRepositoryImpl{repoImpl}
}
