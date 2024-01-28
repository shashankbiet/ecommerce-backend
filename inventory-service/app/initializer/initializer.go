package initializer

import (
	"context"
	"fmt"
	httpserver "inventory-service/app/server/http"
	"inventory-service/pkg/config"
	"inventory-service/pkg/db"
	"inventory-service/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func InitializeConfig() {
	config := config.GetConfig()
	fmt.Printf("config:%+v\n", config)
}

func InitializeLogger() {
	logger.InitLogger()
	logger.Info("logger setup done")
}

func InitializeDb() {
	db.GetSqlConnection()
}

func InitializeServer(ctx context.Context) {

	httpServer, err := httpserver.InitHttpServer()
	if err != nil {
		logger.Error("error in starting http server", "error", err)
	}
	handleGracefulShutdown(ctx, httpServer)
}

func handleGracefulShutdown(ctx context.Context, httpServer *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("error in shutting down http server", "error", err)
	}
}
