package repositoryImpl

import (
	"github.com/tiwariayush700/auction/repository"
	"gorm.io/gorm"
)

type itemRepositoryImpl struct {
	repositoryImpl
}

func NewItemRepositoryImpl(db *gorm.DB) repository.ItemRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &itemRepositoryImpl{repoImpl}
}
