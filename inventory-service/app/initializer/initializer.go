package initializer

import (
	"fmt"
	"inventory-service/app/config"
	"inventory-service/app/datastore"
	"inventory-service/app/handler"
	"inventory-service/app/service"
	"inventory-service/pkg/db"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeConfig() {
	config := config.GetConfig()
	fmt.Printf("config:%+v\n", config)
}

func InitializeDb() {
	db.GetSqlConnection()
}

func InitializeHttp() {
	mySqlCategoryStore := datastore.GetMySqlCategoryStore()
	categoryService := service.NewCategoryService(mySqlCategoryStore)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router := mux.NewRouter()
	router.HandleFunc("/health-check", handler.HealthCheckHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/category", categoryHandler.Add).Methods(http.MethodPost)
	router.HandleFunc("/api/category/{id}", categoryHandler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/category", categoryHandler.GetAll).Methods(http.MethodGet)

	config := config.GetConfig()
	port := fmt.Sprintf(":%d", config.HttpServer.Port)
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Println("HTTP server error: ", err)
	}

}
