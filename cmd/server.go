package cmd

import (
	"fmt"
	"net"

	"github.com/jiro94/elasticsearch-sample/api/interfaces/handler"

	"github.com/jiro94/elasticsearch-sample/api/interfaces/http"
	"github.com/jiro94/elasticsearch-sample/api/registry"
	"github.com/jiro94/elasticsearch-sample/config"
	"github.com/jiro94/elasticsearch-sample/internal/persistence/elasticsearch"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Run:   runServer,
}

var (
	port int
)

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "P", 9000, "API HTTP Server Port")
}

func newListener() (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf(":%d", port))
}

func runServer(_ *cobra.Command, _ []string) {
	cfg := config.Get()

	regOpts := registry.OptionGroup{
		registry.SetElasticsearch(elasticsearch.MustNewClient(cfg.ES)),
	}

	fx.New(
		fx.Provide(
			func() http.Config { return cfg.Server },
			newListener,
			http.NewServer,

			func() registry.OptionGroup { return regOpts },
			registry.New,
		),
		fx.Invoke(
			handler.Register,
		),
	).Run()
}
