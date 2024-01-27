package gin

import (
	"context"
	"errors"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Config struct {
	Mode string `mapstructure:"mode"`
	Port string `mapstructure:"port"`
}

func StartGin(lc fx.Lifecycle, cfg *Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	var e = gin.New()

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: e,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info().Msgf("Starting gin server, listen on %s.", cfg.Port)
			var err error
			go func() {
				if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Msgf("Fail to run gin server, err: %+v", err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			defer log.Info().Msgf("Stopping gin HTTP server.")
			c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			return server.Shutdown(c)
		},
	})

	e.Use(func(c *gin.Context) {
		c.Next()

		if !strings.Contains(c.Request.URL.Path, "playground") && !strings.Contains(c.Request.URL.Path, "health") {
			log.Ctx(c.Request.Context()).Debug().
				Str("url", c.Request.URL.String()).
				Str("method", c.Request.Method).
				Interface("header", c.Request.Header).
				Bool("access_log", true).
				Msgf("access log")
		}
	})
	e.Use(NewRequestIDMiddleware())
	e.Use(NewRecoverMiddleware())
	e.Use(CORS())
	e.Use(Swagger())

	return e
}

// RegisterDefaultRoute provide default handler
func RegisterDefaultRoute(e *gin.Engine) {
	e.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
}