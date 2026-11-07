package httpserver

import (
	"fmt"
	"inventory-service/app/handler"
	"inventory-service/pkg/config"
	"inventory-service/pkg/logger"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitHttpServer() (*http.Server, error) {
	router := mux.NewRouter()
	router.Use(metricsMiddleware)
	router.HandleFunc("/health", handler.HealthCheckHandler).Methods(http.MethodGet)
	router.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
	router.Handle("/debug/pprof", http.HandlerFunc(pprof.Index))
	router.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	router.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	router.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	router.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	router.Handle("/debug/pprof/{cmd}", http.HandlerFunc(pprof.Index))
	if err := registerCategoryRoutes(router); err != nil {
		return nil, err
	}
	if err := registerSubCategoryHandler(router); err != nil {
		return nil, err
	}
	if err := registerProductHandler(router); err != nil {
		return nil, err
	}
	if err := registerInventoryHandler(router); err != nil {
		return nil, err
	}

	port := fmt.Sprintf(":%d", config.GetConfig().InventoryService.Port)
	httpServer := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  time.Duration(config.GetConfig().InventoryService.ReadTimeoutMs) * time.Millisecond,  // maximum duration for reading the entire request
		WriteTimeout: time.Duration(config.GetConfig().InventoryService.WriteTimeoutMs) * time.Millisecond, // maximum duration before timing out writes of the response
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to start http server", "error", err)
			panic(err)
		}
	}()

	return httpServer, nil
}
