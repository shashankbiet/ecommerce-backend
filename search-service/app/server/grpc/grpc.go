package grpcserver

import (
	"context"
	"fmt"
	"net"
	"search-service/app/dao"
	producthandler "search-service/app/handler/product"
	"search-service/app/service"
	"search-service/pkg/config"
	"search-service/pkg/db"
	"search-service/pkg/logger"
	"search-service/pkg/metric"
	searchpb "search-service/proto/search"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
)

func InitGrpcServer(ctx context.Context, ps *metric.PrometheusServer) (*grpc.Server, error) {
	grpcMetrics := ps.CreateGrpcServerMetrics()
	unaryServerInterceptor := []grpc.UnaryServerInterceptor{
		grpcMetrics.UnaryServerInterceptor(),
	}
	streamServerInterceptor := []grpc.StreamServerInterceptor{
		grpcMetrics.StreamServerInterceptor(),
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(unaryServerInterceptor...)),
		grpc.StreamInterceptor(middleware.ChainStreamServer(streamServerInterceptor...)),
	)

	registerServices(ctx, grpcServer)

	port := fmt.Sprintf(":%v", config.GetConfig().GrpcServer.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error("failed to start grpc server", "error", err)
			panic(err)
		}
	}()

	return grpcServer, nil
}

func registerServices(ctx context.Context, grpcServer *grpc.Server) {
	// Register search service
	elasticSearchClient := db.GetElasticSearchConnection()
	indexName := config.GetConfig().ElasticSearch.IndexName
	if indexName == "" {
		panic("ElasticSearch index name is not set")
	}
	datastore := dao.NewProductDataStore(elasticSearchClient, indexName)
	productService := service.NewProductService(datastore)
	searchpb.RegisterSearchServiceServer(grpcServer, producthandler.NewProductServer(productService))
}
