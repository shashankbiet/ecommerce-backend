package initializer

import (
	"context"
	httpserver "inventory-service/app/server/http"
	"inventory-service/pkg/config"
	"inventory-service/pkg/db"
	"inventory-service/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitializeConfig() {
	config.GetConfig()
}

func InitializeLogger() {
	logger.InitLogger()
}

func InitializeDb() error {
	_, err := db.GetSqlConnection()
	return err
}

func InitializeServer(ctx context.Context) error {

	httpServer, err := httpserver.InitHttpServer()
	if err != nil {
		return err
	}
	handleGracefulShutdown(ctx, httpServer)
	return nil
}

func handleGracefulShutdown(ctx context.Context, httpServer *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("error in shutting down http server", "error", err)
	}
}
