package repositoryImpl

import (
	"context"
	userError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/repository"
	"gorm.io/gorm"
)

type itemRepositoryImpl struct {
	repositoryImpl
}

func (i *itemRepositoryImpl) GetItemsByStatus(ctx context.Context, status models.ItemStatus) ([]models.Item, error) {

	items := make([]models.Item, 0)
	err := i.DB.Where("status = ?", string(status)).Find(&items).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, userError.ErrorNotFound
		}
		return nil, err
	}

	return items, nil
}

func (i *itemRepositoryImpl) FetchItems(ctx context.Context) ([]models.Item, error) {
	items := make([]models.Item, 0)
	err := i.DB.Find(&items).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, userError.ErrorNotFound
		}
		return nil, err
	}

	return items, nil
}

func NewItemRepositoryImpl(db *gorm.DB) repository.ItemRepository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &itemRepositoryImpl{repoImpl}
}
