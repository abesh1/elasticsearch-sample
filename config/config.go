package config

//go:generate statik -f -src ./yaml

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jiro94/elasticsearch-sample/api/interfaces/http"
	"github.com/jiro94/elasticsearch-sample/internal/persistence/elasticsearch"
	"github.com/rakyll/statik/fs"
	"go.uber.org/config"

	// statik
	_ "github.com/jiro94/elasticsearch-sample/config/statik"
)

var cfg = &Config{}

type Config struct {
	Env string

	Server http.Config
	ES     elasticsearch.Config
}

func Load(envStr string, overridePath string) error {
	confFs, err := fs.New()
	if err != nil {
		return err
	}

	f, err := confFs.Open(fmt.Sprintf("/config.%s.yml", envStr))
	if err != nil {
		return err
	}
	defer f.Close()

	opts := []config.YAMLOption{config.Source(f)}
	if overridePath != "" {
		overrideAbsPath, err := filepath.Abs(overridePath)
		if err != nil {
			return err
		}
		opts = append(opts, config.File(overrideAbsPath))
	}
	provider, err := config.NewYAML(opts...)
	if err != nil {
		return err
	}

	if err := provider.Get("server").Populate(&cfg.Server); err != nil {
		return err
	}
	if err := provider.Get("elasticsearch").Populate(&cfg.ES); err != nil {
		return err
	}

	cfg.Env = envStr

	return nil
}

func Get() *Config {
	return cfg
}

func Show() error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	return encoder.Encode(cfg)
}
