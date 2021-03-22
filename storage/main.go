package storage

import (
	"github.com/Yangiboev/golang-mongodb-monolith/storage/mongo"
	"github.com/Yangiboev/golang-mongodb-monolith/storage/repo"
	db "go.mongodb.org/mongo-driver/mongo"
)

type StorageI interface {
	Product() repo.ProductStorageI
}

type product struct {
	productRepo repo.ProductStorageI
}

func NewProductStorage(db *db.Database) StorageI {
	return &product{
		productRepo: mongo.NewProductRepo(db),
	}
}

func (pr product) Product() repo.ProductStorageI {
	return pr.productRepo
}
