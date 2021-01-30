package repository

import (
	"context"
	"github.com/tiwariayush700/auction/models"
)

type ItemRepository interface {
	Repository
	GetItemsByStatus(ctx context.Context, status models.ItemStatus) ([]models.Item, error)
	FetchItems(ctx context.Context) ([]models.Item, error)
}
