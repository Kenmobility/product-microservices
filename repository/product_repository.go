package repository

import (
	"github.com/kenmobility/product-microservice/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProductByPublicID(id string) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(id string) error
	ListProducts(filter string) ([]*models.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) CreateProduct(product *models.Product) (*models.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepo) GetProductByPublicID(id string) (*models.Product, error) {
	var product models.Product
	if err := r.db.Preload("DigitalProduct").Preload("PhysicalProduct").Preload("SubscriptionProduct").
		Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) UpdateProduct(product *models.Product) (*models.Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepo) DeleteProduct(id string) error {
	if err := r.db.Delete(&models.Product{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *productRepo) ListProducts(filter string) ([]*models.Product, error) {
	var products []*models.Product
	query := r.db.Preload("DigitalProduct").Preload("PhysicalProduct").Preload("SubscriptionProduct")

	if filter != "" {
		query = query.Where("type = ?", filter)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
