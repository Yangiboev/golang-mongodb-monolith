package repo

import (
	"github.com/Yangiboev/golang-mongodb-monolith/api/models"
)

type ProductStorageI interface {
	Create(product *models.Product) (string, error)
	Get(id string) (*models.Product, error)
	GetAll(page, limit int64, name string) ([]*models.Product, int64, error)
}
