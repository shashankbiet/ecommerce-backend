package initializer

import (
	"fmt"
	"inventory-service/app/config"
	"inventory-service/app/handler"
	"inventory-service/app/utils"
	"inventory-service/pkg/db"
	"inventory-service/pkg/logger"
	"net/http"

	consoleLog "inventory-service/pkg/logger/console"

	"github.com/gorilla/mux"
)

func InitializeConfig() {
	config := config.GetConfig()
	fmt.Printf("config:%+v\n", config)
}

func InitializeLogger() {
	logLevel := utils.GetLogLevel()
	log := consoleLog.GetConsoleLog()
	logger.InitLogger(log, logLevel)
	logger.Log.Info("logger setup done")
}

func InitializeDb() {
	db.GetSqlConnection()
}

func InitializeHttp() {
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.HealthCheckHandler).Methods(http.MethodGet)
	registerCategoryRoutes(router)
	registerSubCategoryHandler(router)
	registerProductHandler(router)
	registerInventoryHandler(router)

	config := config.GetConfig()
	port := fmt.Sprintf(":%d", config.HttpServer.Port)
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Log.Info("http server error", "error", err)
	}
}
