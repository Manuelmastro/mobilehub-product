package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/Manuelmastro/mobilehub-product/pkg/db"
	"github.com/Manuelmastro/mobilehub-product/pkg/models"
	"github.com/Manuelmastro/mobilehub-product/pkg/pb"
)

type ProductServiceServer struct {
	pb.UnimplementedProductServiceServer
	H db.Handler
}

func (s *ProductServiceServer) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	var products []models.Product
	if err := s.H.DB.Where("deleted_at IS NULL").Find(&products).Error; err != nil {
		return nil, errors.New("failed to fetch products")
	}
	var response []*pb.Product
	for _, product := range products {
		response = append(response, &pb.Product{
			Id:           fmt.Sprint(product.ID),
			ProductName:  product.ProductName,
			Description:  product.Description,
			ImageUrl:     product.ImageUrl,
			Price:        float32(product.Price),
			Stock:        product.Stock,
			CategoryName: product.CategoryName,
		})
	}

	return &pb.GetProductsResponse{Products: response}, nil
}

func (s *ProductServiceServer) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	// Authorization check for admin role
	role, ok := ctx.Value("role").(string)
	if !ok || role != "admin" {
		return nil, errors.New("unauthorized: only admin can add products")
	}

	product := models.Product{
		ProductName:  req.ProductName,
		Description:  req.Description,
		ImageUrl:     req.ImageUrl,
		Price:        float64(req.Price),
		Stock:        req.Stock,
		CategoryName: req.CategoryName,
	}

	if err := s.H.DB.Create(&product).Error; err != nil {
		return nil, errors.New("failed to add product")
	}

	return &pb.AddProductResponse{
		Message: "Product added successfully",
	}, nil
}

func (s *ProductServiceServer) EditProduct(ctx context.Context, req *pb.EditProductRequest) (*pb.EditProductResponse, error) {
	// Authorization check for admin role
	role, ok := ctx.Value("role").(string)
	if !ok || role != "admin" {
		return nil, errors.New("unauthorized: only admin can edit products")
	}

	productID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	var product models.Product
	if err := s.H.DB.First(&product, productID).Error; err != nil {
		return nil, errors.New("product not found")
	}

	product.ProductName = req.ProductName
	product.Description = req.Description
	product.ImageUrl = req.ImageUrl
	product.Price = float64(req.Price)
	product.Stock = req.Stock
	product.CategoryName = req.CategoryName

	if err := s.H.DB.Save(&product).Error; err != nil {
		return nil, errors.New("failed to update product")
	}

	return &pb.EditProductResponse{
		Message: "Product updated successfully",
	}, nil
}

func (s *ProductServiceServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	// Authorization check for admin role
	role, ok := ctx.Value("role").(string)
	if !ok || role != "admin" {
		return nil, errors.New("unauthorized: only admin can delete products")
	}

	productID, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	var product models.Product
	if err := s.H.DB.First(&product, productID).Error; err != nil {
		return nil, errors.New("product not found")
	}

	if err := s.H.DB.Delete(&product).Error; err != nil {
		return nil, errors.New("failed to delete product")
	}

	return &pb.DeleteProductResponse{
		Message: "Product deleted successfully",
	}, nil
}

func (s *ProductServiceServer) ViewProducts(ctx context.Context, req *pb.ViewProductsRequest) (*pb.ViewProductsResponse, error) {
	// No authentication needed for viewing products
	var products []models.Product
	if err := s.H.DB.Where("deleted_at IS NULL").Find(&products).Error; err != nil {
		return nil, errors.New("failed to fetch products")
	}

	var response []*pb.Product
	for _, product := range products {
		response = append(response, &pb.Product{
			Id:           fmt.Sprint(product.ID),
			ProductName:  product.ProductName,
			Description:  product.Description,
			ImageUrl:     product.ImageUrl,
			Price:        float32(product.Price),
			Stock:        product.Stock,
			CategoryName: product.CategoryName,
		})
	}

	return &pb.ViewProductsResponse{Products: response}, nil
}

func (s *ProductServiceServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var product models.Product

	// Fetch product by ID
	if err := s.H.DB.Where("id = ?", req.Id).First(&product).Error; err != nil {
		return nil, errors.New("product not found")
	}

	// Map product to response
	response := &pb.GetProductResponse{
		Product: &pb.Product{
			Id:           fmt.Sprint(product.ID),
			ProductName:  product.ProductName,
			Description:  product.Description,
			ImageUrl:     product.ImageUrl,
			Price:        float32(product.Price),
			Stock:        int32(product.Stock),
			CategoryName: product.CategoryName,
		},
	}

	return response, nil
}
