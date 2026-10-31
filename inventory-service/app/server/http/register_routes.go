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
	router.HandleFunc("/api/inventory/{productId}", inventoryHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/inventory", inventoryHandler.GetAll).Methods(http.MethodGet)
	return nil
}
