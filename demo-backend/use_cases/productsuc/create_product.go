package productsuc

import (
	"context"
	"demo-backend/domain/entities"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name  string  `validate:"required,min=3,max=15"`
	Price float32 `validate:"required"`
}

type CreateProductResponse struct {
	Product *entities.Product
}

type createProductUseCase struct {
	dataStore ProductsDataStorer
}

func NewCreateProductUseCase(ds ProductsDataStorer) *createProductUseCase {
	return &createProductUseCase{
		dataStore: ds,
	}
}

func (uc *createProductUseCase) CreateProduct(ctx context.Context, request CreateProductRequest) (*CreateProductResponse, error) {

	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}

	var product = &entities.Product{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      request.Name,
		Price:     request.Price,
	}

	err = uc.dataStore.Create(product)
	if err != nil {
		return nil, err
	}

	var response = &CreateProductResponse{Product: product}
	return response, nil
}
