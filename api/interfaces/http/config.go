package http

import "time"

type Config struct {
	ReadTimeout       time.Duration `yaml:"read_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"`
	EnableKeepAlive   bool          `yaml:"enable_keep_alive"`
	MaxConns          int           `yaml:"max_conns"`
	Concurrency       int
}
