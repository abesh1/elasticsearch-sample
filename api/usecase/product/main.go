package product

import (
	"github.com/jiro94/elasticsearch-sample/api/domain/repository"
	"github.com/jiro94/elasticsearch-sample/api/domain/service"
)

func NewService(repo repository.Product) service.Product {
	return &serv{repo: repo}
}

type serv struct {
	repo repository.Product
}
