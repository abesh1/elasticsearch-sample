package service

import (
	"context"

	"github.com/jiro94/elasticsearch-sample/api/domain/entityreq"
)

type Product interface {
	GetSearch(context.Context, entityreq.ProductSearch) (interface{}, error)
	GetSearchSuggestion(context.Context, entityreq.ProductSearch) (interface{}, error)
	InsertSearchSeed() error
}
