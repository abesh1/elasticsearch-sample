package http

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
	"golang.org/x/net/netutil"
)

func NewServer(lc fx.Lifecycle, cfg Config, ln net.Listener) (*http.Server, error) {
	server := &http.Server{
		ReadTimeout:       cfg.ReadTimeout * time.Millisecond,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout * time.Millisecond,
		WriteTimeout:      cfg.WriteTimeout * time.Millisecond,
		IdleTimeout:       cfg.IdleTimeout * time.Millisecond,
		ErrorLog:          log.New(os.Stderr, "[http server error] ", log.LstdFlags|log.Lshortfile),
	}
	server.SetKeepAlivesEnabled(cfg.EnableKeepAlive)

	if cfg.MaxConns > 0 {
		ln = netutil.LimitListener(ln, cfg.MaxConns)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return server, nil
}
