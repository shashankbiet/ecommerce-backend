package httpserver

import (
	"fmt"
	"inventory-service/app/handler"
	"inventory-service/pkg/config"
	"inventory-service/pkg/logger"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func InitHttpServer() (*http.Server, error) {
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.HealthCheckHandler).Methods(http.MethodGet)
	registerCategoryRoutes(router)
	registerSubCategoryHandler(router)
	registerProductHandler(router)
	registerInventoryHandler(router)

	port := fmt.Sprintf(":%d", config.GetConfig().HttpServer.Port)
	httpServer := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  time.Duration(config.GetConfig().HttpServer.ReadTimeoutMs) * time.Millisecond,  // maximum duration for reading the entire request
		WriteTimeout: time.Duration(config.GetConfig().HttpServer.WriteTimeoutMs) * time.Millisecond, // maximum duration before timing out writes of the response
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Error("failed to start http server", "error", err)
			panic(err)
		}
	}()

	return httpServer, nil
}
