package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jiro94/elasticsearch-sample/api/domain/entityreq"
	"github.com/jiro94/elasticsearch-sample/api/domain/service"
)

func getSearchEndpoint(s service.Product) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetSearch(ctx, request.(entityreq.ProductSearch))
	}
}

func getSearchSuggestionEndpoint(s service.Product) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetSearchSuggestion(ctx, request.(entityreq.ProductSearch))
	}
}
