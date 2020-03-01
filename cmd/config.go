package cmd

import (
	"github.com/jiro94/elasticsearch-sample/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show config",
	RunE:  runConfig,
}

func runConfig(_ *cobra.Command, _ []string) error {
	return config.Show()
}

func loadConfig(_ *cobra.Command, _ []string) error {
	return config.Load(env, confPath)
}
