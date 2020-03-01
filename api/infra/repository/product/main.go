package product

import (
	"github.com/jiro94/elasticsearch-sample/api/domain/repository"
	"github.com/jiro94/elasticsearch-sample/internal/persistence/elasticsearch"
)

func New(es elasticsearch.Client) repository.Product {
	return &repo{
		ES: es,
	}
}

type repo struct {
	ES elasticsearch.Client
}
