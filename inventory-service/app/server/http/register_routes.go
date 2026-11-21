package httpserver

import (
	"fmt"
	pgDao "inventory-service/app/dao/pg"
	"inventory-service/app/handler"
	"inventory-service/app/producer/kafka"
	"inventory-service/app/service"
	"net/http"

	"github.com/gorilla/mux"
)

func registerBrandHandler(router *mux.Router) error {
	brandStore, err := pgDao.GetPgBrandStore()
	if err != nil {
		return fmt.Errorf("initialize PostgreSQL brand store: %w", err)
	}
	brandService := service.NewBrandService(brandStore)
	brandHandler := handler.NewBrandHandler(brandService)

	router.HandleFunc("/api/brand", brandHandler.GetAll).Methods(http.MethodGet)
	return nil
}

func registerCategoryRoutes(router *mux.Router) error {
	categoryStore, err := pgDao.GetPgCategoryStore()
	if err != nil {
		return fmt.Errorf("initialize PostgreSQL category store: %w", err)
	}
	categoryService := service.NewCategoryService(categoryStore)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router.HandleFunc("/api/category", categoryHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/category/{id}", categoryHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/category", categoryHandler.GetAll).Methods(http.MethodGet)
	return nil
}

func registerSubCategoryHandler(router *mux.Router) error {
	subCategoryStore, err := pgDao.GetPgSubCategoryStore()
	if err != nil {
		return fmt.Errorf("initialize PostgreSQL subcategory store: %w", err)
	}
	categoryService := service.NewSubCategoryService(subCategoryStore)
	subCategoryHandler := handler.NewSubCategoryHandler(categoryService)

	router.HandleFunc("/api/subcategory", subCategoryHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/subcategory/{id}", subCategoryHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/subcategory", subCategoryHandler.GetAll).Methods(http.MethodGet)
	return nil
}

func registerProductHandler(router *mux.Router) error {
	productStore, err := pgDao.GetPgProductStore()
	if err != nil {
		return fmt.Errorf("initialize PostgreSQL product store: %w", err)
	}
	kafkaProductProducer := kafka.GetKafkaProductProducer()
	productService := service.NewProductService(productStore, kafkaProductProducer)
	productHandler := handler.NewProductHandler(productService)

	router.HandleFunc("/api/product", productHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/product", productHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/product/{id}", productHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/product", productHandler.GetAll).Methods(http.MethodGet)
	return nil
}

func registerInventoryHandler(router *mux.Router) error {
	inventoryStore, err := pgDao.GetPgInventoryStore()
	if err != nil {
		return fmt.Errorf("initialize PostgreSQL inventory store: %w", err)
	}
	kafkaInventoryProducer := kafka.GetKafkaInventoryProducer()
	inventoryService := service.NewInventoryService(inventoryStore, kafkaInventoryProducer)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)

	router.HandleFunc("/api/inventory", inventoryHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/inventory", inventoryHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/inventory/{product_id}", inventoryHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/inventory", inventoryHandler.GetAll).Methods(http.MethodGet)
	return nil
}

func registerKeywordHandler(router *mux.Router) error {
	keywordStore, err := pgDao.GetPgKeywordStore()
	if err != nil {
		return fmt.Errorf("initialize PostgreSQL keyword store: %w", err)
	}
	keywordService := service.NewKeywordService(keywordStore)
	keywordHandler := handler.NewKeywordHandler(keywordService)

	router.HandleFunc("/api/keyword", keywordHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/keyword/{id}", keywordHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/keyword", keywordHandler.GetAll).Methods(http.MethodGet)
	return nil
}

func registerProductKeywordHandler(router *mux.Router) error {
	productKeywordStore, err := pgDao.GetPgProductKeywordStore()
	if err != nil {
		return fmt.Errorf("initialize PostgreSQL product keyword store: %w", err)
	}
	productKeywordService := service.NewProductKeywordService(productKeywordStore)
	productKeywordHandler := handler.NewProductKeywordHandler(productKeywordService)

	router.HandleFunc("/api/product-keyword", productKeywordHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/product-keyword/product/{product_id}", productKeywordHandler.GetByProductId).Methods(http.MethodGet)
	router.HandleFunc("/api/product-keyword/keyword/{keyword_id}", productKeywordHandler.GetByKeywordId).Methods(http.MethodGet)
	router.HandleFunc("/api/product-keyword/{product_id}/{keyword_id}", productKeywordHandler.Delete).Methods(http.MethodDelete)
	return nil
}
