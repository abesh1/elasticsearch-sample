package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	PersistentPreRunE: loadConfig,
}

var (
	env      string
	confPath string
)

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(seedCmd)

	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "dev", "Environment")
	rootCmd.PersistentFlags().StringVarP(&confPath, "conf", "c", "", "Load configuration from `FILE`")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(fmt.Sprintf("%+v", err))
	}
}
