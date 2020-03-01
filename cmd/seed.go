package cmd

import (
	"context"

	"github.com/jiro94/elasticsearch-sample/api/registry"
	"github.com/jiro94/elasticsearch-sample/config"
	"github.com/jiro94/elasticsearch-sample/internal/persistence/elasticsearch"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Insert seed",
	Run:   insertSeed,
}

func insertSeed(_ *cobra.Command, _ []string) {
	cfg := config.Get()

	regOpts := registry.OptionGroup{
		registry.SetElasticsearch(elasticsearch.MustNewClient(cfg.ES)),
	}

	fx.New(
		fx.Provide(
			func() registry.OptionGroup { return regOpts },
			registry.New,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, s *registry.Services) error {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return s.Product.InsertSearchSeed(ctx)
					},
				})
				return nil
			},
		),
	).Run()
}
