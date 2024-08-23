package handlers

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/kenmobility/product-microservice/models"
	"github.com/kenmobility/product-microservice/pb"
	"github.com/kenmobility/product-microservice/service"
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := &models.Product{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Price:       req.Product.Price,
	}

	// Set type-specific fields
	switch t := req.Product.ProductType.(type) {
	case *pb.Product_DigitalProduct:
		product.DigitalProduct = &models.DigitalProduct{
			FileSize:     t.DigitalProduct.FileSize,
			DownloadLink: t.DigitalProduct.DownloadLink,
		}
	case *pb.Product_PhysicalProduct:
		product.PhysicalProduct = &models.PhysicalProduct{
			Weight:     t.PhysicalProduct.Weight,
			Dimensions: t.PhysicalProduct.Dimensions,
		}
	case *pb.Product_SubscriptionProduct:
		product.SubscriptionProduct = &models.SubscriptionProduct{
			SubscriptionPeriod: t.SubscriptionProduct.SubscriptionPeriod,
			RenewalPrice:       t.SubscriptionProduct.RenewalPrice,
		}
	}

	createdProduct, err := h.service.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	createdAt, _ := ptypes.TimestampProto(createdProduct.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(createdProduct.UpdatedAt)

	response := &pb.CreateProductResponse{
		Product: &pb.Product{
			PublicId:    createdProduct.PublicID,
			Name:        createdProduct.Name,
			Description: createdProduct.Description,
			Price:       createdProduct.Price,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		},
	}

	return response, nil
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := h.service.GetProductByID(req.Id)
	if err != nil {
		return nil, err
	}

	createdAt, _ := ptypes.TimestampProto(product.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(product.UpdatedAt)

	protoProduct := &pb.Product{
		Id:          int32(product.ID),
		PublicId:    product.PublicID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	// Set type-specific fields in the response
	switch {
	case product.DigitalProduct != nil:
		protoProduct.ProductType = &pb.Product_DigitalProduct{
			DigitalProduct: &pb.DigitalProduct{
				FileSize:     product.DigitalProduct.FileSize,
				DownloadLink: product.DigitalProduct.DownloadLink,
			},
		}
	case product.PhysicalProduct != nil:
		protoProduct.ProductType = &pb.Product_PhysicalProduct{
			PhysicalProduct: &pb.PhysicalProduct{
				Weight:     product.PhysicalProduct.Weight,
				Dimensions: product.PhysicalProduct.Dimensions,
			},
		}
	case product.SubscriptionProduct != nil:
		protoProduct.ProductType = &pb.Product_SubscriptionProduct{
			SubscriptionProduct: &pb.SubscriptionProduct{
				SubscriptionPeriod: product.SubscriptionProduct.SubscriptionPeriod,
				RenewalPrice:       product.SubscriptionProduct.RenewalPrice,
			},
		}
	}

	response := &pb.GetProductResponse{
		Product: protoProduct,
	}

	return response, nil
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {

	product := &models.Product{
		Name:        req.Product.Name,
		Description: req.Product.Description,
		Price:       req.Product.Price,
	}

	switch t := req.Product.ProductType.(type) {
	case *pb.Product_DigitalProduct:
		product.DigitalProduct = &models.DigitalProduct{
			ProductID:    uint(req.Product.Id),
			FileSize:     t.DigitalProduct.FileSize,
			DownloadLink: t.DigitalProduct.DownloadLink,
		}
	case *pb.Product_PhysicalProduct:
		product.PhysicalProduct = &models.PhysicalProduct{
			ProductID:  uint(req.Product.Id),
			Weight:     t.PhysicalProduct.Weight,
			Dimensions: t.PhysicalProduct.Dimensions,
		}
	case *pb.Product_SubscriptionProduct:
		product.SubscriptionProduct = &models.SubscriptionProduct{
			ProductID:          uint(req.Product.Id),
			SubscriptionPeriod: t.SubscriptionProduct.SubscriptionPeriod,
			RenewalPrice:       t.SubscriptionProduct.RenewalPrice,
		}
	}

	updatedProduct, err := h.service.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	createdAt, _ := ptypes.TimestampProto(updatedProduct.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(updatedProduct.UpdatedAt)

	response := &pb.UpdateProductResponse{
		Product: &pb.Product{
			Id:          int32(updatedProduct.ID),
			PublicId:    updatedProduct.PublicID,
			Name:        updatedProduct.Name,
			Description: updatedProduct.Description,
			Price:       updatedProduct.Price,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		},
	}

	return response, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	if err := h.service.DeleteProduct(req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteProductResponse{}, nil
}

func (h *ProductHandler) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := h.service.ListProducts(req.TypeFilter)
	if err != nil {
		return nil, err
	}

	var protoProducts []*pb.Product
	for _, p := range products {
		createdAt, _ := ptypes.TimestampProto(p.CreatedAt)
		updatedAt, _ := ptypes.TimestampProto(p.UpdatedAt)

		protoProduct := &pb.Product{
			Id:          int32(p.ID),
			PublicId:    p.PublicID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}

		switch {
		case p.DigitalProduct != nil:
			protoProduct.ProductType = &pb.Product_DigitalProduct{
				DigitalProduct: &pb.DigitalProduct{
					FileSize:     p.DigitalProduct.FileSize,
					DownloadLink: p.DigitalProduct.DownloadLink,
				},
			}
		case p.PhysicalProduct != nil:
			protoProduct.ProductType = &pb.Product_PhysicalProduct{
				PhysicalProduct: &pb.PhysicalProduct{
					Weight:     p.PhysicalProduct.Weight,
					Dimensions: p.PhysicalProduct.Dimensions,
				},
			}
		case p.SubscriptionProduct != nil:
			protoProduct.ProductType = &pb.Product_SubscriptionProduct{
				SubscriptionProduct: &pb.SubscriptionProduct{
					SubscriptionPeriod: p.SubscriptionProduct.SubscriptionPeriod,
					RenewalPrice:       p.SubscriptionProduct.RenewalPrice,
				},
			}
		}

		protoProducts = append(protoProducts, protoProduct)
	}

	return &pb.ListProductsResponse{Products: protoProducts}, nil
}
