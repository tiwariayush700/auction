package serviceImpl

import (
	"context"
	"github.com/tiwariayush700/auction/models"
	"github.com/tiwariayush700/auction/repository"
	"github.com/tiwariayush700/auction/services"
)

type itemServiceImpl struct {
	repository repository.ItemRepository
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
