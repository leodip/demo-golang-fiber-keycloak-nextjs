package datastores

import (
	"demo-backend/domain/entities"
	"sort"
	"sync"

	"github.com/google/uuid"
)

type productsDataStore struct {
	products map[uuid.UUID]entities.Product
	sync.Mutex
}

func NewProductsDataStore() *productsDataStore {
	return &productsDataStore{
		products: make(map[uuid.UUID]entities.Product),
	}
}

func (ds *productsDataStore) Create(product *entities.Product) error {
	ds.Lock()
	ds.products[product.Id] = *product
	ds.Unlock()
	return nil
}

func (ds *productsDataStore) GetAll() []entities.Product {
	all := make([]entities.Product, 0, len(ds.products))
	for _, value := range ds.products {
		all = append(all, value)
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].CreatedAt.Before(all[j].CreatedAt)
	})
	return all
}
