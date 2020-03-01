package registry

import "github.com/jiro94/elasticsearch-sample/internal/persistence/elasticsearch"

type (
	Option      func(*registry)
	OptionGroup []Option
)

func SetElasticsearch(v elasticsearch.Client) Option {
	return func(r *registry) {
		r.es = v
	}
}
