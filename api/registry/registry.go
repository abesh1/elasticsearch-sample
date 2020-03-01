package registry

import (
	"context"

	"github.com/jiro94/elasticsearch-sample/api/domain/service"

	"github.com/jiro94/elasticsearch-sample/internal/persistence/elasticsearch"
	"go.uber.org/fx"
)

func New(lc fx.Lifecycle, opts OptionGroup) *Services {
	reg := &registry{}
	for _, o := range opts {
		o(reg)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			reg.closeRepositoryConn()
			return nil
		},
	})

	return &Services{Product: reg.newProductService()}
}

type registry struct {
	es elasticsearch.Client
}

func (r registry) closeRepositoryConn() {
	if r.es != nil {
		r.es.Close()
	}
}

type Services struct {
	Product service.Product
}
