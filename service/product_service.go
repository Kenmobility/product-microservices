package service

import (
	"github.com/google/uuid"
	"github.com/kenmobility/product-microservice/models"
	"github.com/kenmobility/product-microservice/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	product.PublicID = uuid.NewString()
	return s.repo.CreateProduct(product)
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.repo.GetProductByPublicID(id)
}

func (s *ProductService) UpdateProduct(product *models.Product) (*models.Product, error) {
	return s.repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.DeleteProduct(id)
}

func (s *ProductService) ListProducts(filter string) ([]*models.Product, error) {
	return s.repo.ListProducts(filter)
}
