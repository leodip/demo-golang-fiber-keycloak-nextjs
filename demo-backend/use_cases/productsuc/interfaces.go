package productsuc

import (
	"demo-backend/domain/entities"
)

type ProductsDataStorer interface {
	GetAll() []entities.Product
	Create(product *entities.Product) error
}
