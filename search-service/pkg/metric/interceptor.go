package metric

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryMetricsInterceptor updates the custom search_service metrics in store.go for each unary RPC.
func UnaryMetricsInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)

		RequestCounter.WithLabelValues("grpc", info.FullMethod).Inc()
		RequestLatency.WithLabelValues(clientName(ctx), "grpc", info.FullMethod).Observe(time.Since(start).Seconds())
		if err != nil {
			code := status.Code(err)
			ErrorCounter.WithLabelValues("grpc", info.FullMethod, code.String()).Inc()
			if code == codes.DeadlineExceeded {
				TimeoutCounter.WithLabelValues(info.FullMethod).Inc()
			}
		}
		return resp, err
	}
}

// clientName reads the optional "client-name" metadata header and falls back to "unknown".
func clientName(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if values := md.Get("client-name"); len(values) > 0 && values[0] != "" {
			return values[0]
		}
	}
	return "unknown"
}
