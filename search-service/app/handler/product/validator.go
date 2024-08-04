package producthandler

import (
	searchpb "search-service/proto/search"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateGetProductRequest(req *searchpb.ProductSearchRequest) error {
	if req.GetCategory() == "" {
		return status.Errorf(codes.InvalidArgument, "Bad Request!")
	}
	return nil
}
