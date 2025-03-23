package httpserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("http_server",
	fx.Provide(NewServer),
	fx.Invoke(StartHttpServer),
)

type Config struct {
	Addr         string
	Handler      http.Handler
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	TlsConfig    *tls.Config
}

func NewServer(config *Config) *http.Server {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Addr),
		Handler:      config.Handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	if config.TlsConfig != nil {
		srv.TLSConfig = config.TlsConfig
	}

	return srv
}

func StartHttpServer(lc fx.Lifecycle, srv *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			zap.L().Info("starting application")
			lis, err := tls.Listen("tcp", srv.Addr, srv.TLSConfig)
			if err != nil {
				return err
			}

			err = srv.Serve(lis)

			return
		},
		OnStop: func(ctx context.Context) error {
			zap.L().Info("application shutdown")
			return srv.Shutdown(ctx)
		},
	})
}
