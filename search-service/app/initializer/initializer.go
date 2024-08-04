package initializer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"search-service/app/consumer"
	grpcserver "search-service/app/server/grpc"
	"search-service/pkg/config"
	"search-service/pkg/db"
	"search-service/pkg/logger"
	"search-service/pkg/metric"
	"syscall"

	"google.golang.org/grpc"
)

func InitializePrometheusServer() *metric.PrometheusServer {
	prometheusServer := &metric.PrometheusServer{}
	prometheusServer.NewPrometheusServer("ecommerce_backend", "search_service")
	metric.InitMetricStore(prometheusServer)
	return prometheusServer
}

func InitializerConfig() {
	config := config.GetConfig()
	fmt.Printf("config:%+v\n", config)
}

func InitializeLogger() {
	logger.InitLogger()
	logger.Info("logger setup done")
}

func InitializeDb() {
	db.GetElasticSearchConnection()
}

func InitializeProductConsumer() {
	c := consumer.NewProductConsumer()
	c.Consume()
}

func InitializeServer(ctx context.Context, ps *metric.PrometheusServer) {
	grpcServer, err := grpcserver.InitGrpcServer(ctx, ps)
	if err != nil {
		logger.Error("error in starting gRPC server", "error", err)
	}
	handleGracefulShutdown(ctx, grpcServer)
}

func handleGracefulShutdown(ctx context.Context, grpcServer *grpc.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	grpcServer.GracefulStop()
}
