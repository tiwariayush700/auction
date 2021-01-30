package services

import (
	"context"
	"github.com/tiwariayush700/auction/models"
)

type ItemService interface {
	CreateItem(ctx context.Context, item *models.Item) error
	GetItemsForStatus(ctx context.Context, status models.ItemStatus) ([]models.Item, error)
	GetItemByID(ctx context.Context, itemID uint) (*models.Item, error)
}
