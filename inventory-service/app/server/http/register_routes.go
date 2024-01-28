package httpserver

import (
	mysqlDao "inventory-service/app/dao/mysql"
	"inventory-service/app/handler"
	"inventory-service/app/producer/kafka"
	"inventory-service/app/service"
	"net/http"

	"github.com/gorilla/mux"
)

func registerCategoryRoutes(router *mux.Router) {
	mySqlCategoryStore := mysqlDao.GetMySqlCategoryStore()
	categoryService := service.NewCategoryService(mySqlCategoryStore)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router.HandleFunc("/api/category", categoryHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/category/{id}", categoryHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/category", categoryHandler.GetAll).Methods(http.MethodGet)
}

func registerSubCategoryHandler(router *mux.Router) {
	mySqlSubCategoryStore := mysqlDao.GetMySqlSubCategoryStore()
	categoryService := service.NewSubCategoryService(mySqlSubCategoryStore)
	subCategoryHandler := handler.NewSubCategoryHandler(categoryService)

	router.HandleFunc("/api/subcategory", subCategoryHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/subcategory/{id}", subCategoryHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/subcategory", subCategoryHandler.GetAll).Methods(http.MethodGet)
}

func registerProductHandler(router *mux.Router) {
	mySqlProductStore := mysqlDao.GetMySqlProductStore()
	kafkaProductProducer := kafka.GetKafkaProductProducer()
	productService := service.NewProductService(mySqlProductStore, kafkaProductProducer)
	productHandler := handler.NewProductHandler(productService)

	router.HandleFunc("/api/product", productHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/product", productHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/product/{id}", productHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/product", productHandler.GetAll).Methods(http.MethodGet)
}

func registerInventoryHandler(router *mux.Router) {
	mySqlInventoryStore := mysqlDao.GetMySqlInventoryStore()
	kafkaInventoryProducer := kafka.GetKafkaInventoryProducer()
	inventoryService := service.NewInventoryService(mySqlInventoryStore, kafkaInventoryProducer)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)

	router.HandleFunc("/api/inventory", inventoryHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/inventory", inventoryHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/inventory/{productId}", inventoryHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/inventory", inventoryHandler.GetAll).Methods(http.MethodGet)
}
