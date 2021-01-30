package serviceImpl

import (
	"github.com/tiwariayush700/auction/repository"
	"github.com/tiwariayush700/auction/services"
)

type itemServiceImpl struct {
	repository repository.ItemRepository
}

func NewItemServiceImpl(repository repository.ItemRepository) services.ItemService {
	return &itemServiceImpl{repository: repository}
}
