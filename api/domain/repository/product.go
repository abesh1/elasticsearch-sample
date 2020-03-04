package repository

import (
	"context"

	"github.com/jiro94/elasticsearch-sample/api/domain/entity"

	"github.com/jiro94/elasticsearch-sample/api/domain/entityreq"
)

type Product interface {
	GetWebESSearch(context.Context, entityreq.ProductSearch) (entity.ESProductList, error)
	GetV5PrefixESSearch(context.Context, entityreq.ProductSearch) (entity.ESProductList, error)
	GetV5PartialESSearch(context.Context, entityreq.ProductSearch) (entity.ESProductList, error)
	GetWebESSearchSuggestion(context.Context, entityreq.ProductSearch) (entity.MultiESProductList, error)
	GetV5PrefixESSearchSuggestion(context.Context, entityreq.ProductSearch) (entity.MultiESObjectList, error)
	GetV5PartialESSearchSuggestion(context.Context, entityreq.ProductSearch) (entity.MultiESObjectList, error)
	GetV5PrefixAndPartialESSearchSuggestion(context.Context, entityreq.ProductSearch) (entity.MultiESObjectList, error)
	InsertSearchSeed(context.Context, entity.ESProductList) error
	InsertSearchAuthorSeed(context.Context, entity.ESObjectList) error
}
