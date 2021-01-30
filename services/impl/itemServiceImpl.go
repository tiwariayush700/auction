package serviceImpl

import (
	"context"
	auctionError "github.com/tiwariayush700/auction/error"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/repository"
	"github.com/tiwariayush700/auction/services"
	"gorm.io/gorm"
)

type itemServiceImpl struct {
	repository repository.ItemRepository
}

func (i *itemServiceImpl) GetItemByID(ctx context.Context, itemID uint) (*models.Item, error) {

	item := &models.Item{}
	err := i.repository.Get(ctx, item, itemID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, auctionError.ErrorNotFound
		}
		return nil, err
	}

	return item, err
}

func (i *itemServiceImpl) CreateItem(ctx context.Context, item *models.Item) error {
	return i.repository.Create(ctx, item)
}

func (i *itemServiceImpl) GetItemsForStatus(ctx context.Context, status models.ItemStatus) ([]models.Item, error) {

	if len(status) > 0 {
		return i.repository.GetItemsByStatus(ctx, status)
	}

	return i.repository.FetchItems(ctx)
}

func NewItemServiceImpl(repository repository.ItemRepository) services.ItemService {
	return &itemServiceImpl{repository: repository}
}
