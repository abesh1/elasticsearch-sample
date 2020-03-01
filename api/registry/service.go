package registry

import (
	"github.com/jiro94/elasticsearch-sample/api/domain/service"
	productrepo "github.com/jiro94/elasticsearch-sample/api/infra/repository/product"
	"github.com/jiro94/elasticsearch-sample/api/usecase/product"
)

func (r registry) newProductService() service.Product {
	return product.NewService(productrepo.New(r.es))
}
