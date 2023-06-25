package logger

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/robinjoseph08/golib/logger"
)

const (
	levelHeader   = "x-log-level"
	versionHeader = "x-version"
	echoIDKey     = "id"
)

// Middleware attaches a logger.Logger instance with a request ID onto the context. It
// also logs every request along with metadata about the request.
func Middleware() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Create logger.
			level := c.Request().Header.Get(levelHeader)
			l := logger.NewWithLevel(level)

			// Save start time of request.
			t1 := time.Now()

			// Set ID.
			id, err := uuid.NewRandom()
			if err != nil {
				return errors.WithStack(err)
			}
			idStr := id.String()
			c.Set(echoIDKey, idStr)

			// Get frontend version.
			version := c.Request().Header.Get(versionHeader)
			if version == "" {
				version = "unknown"
			}

			// Add relevant context to the logger about the request.
			log := l.ID(idStr).Root(logger.Data{
				"method":  c.Request().Method,
				"path":    c.Request().URL.Path,
				"route":   c.Path(),
				"version": version,
			})

			// We set the logger on the underlying context.Context instead of
			// the echo.Context so that if we need to use the underlying
			// context.Context for anything during the request lifecycle, it
			// will also have the request ID on it.
			c.SetRequest(c.Request().WithContext(log.WithContext(c.Request().Context())))
			if err := next(c); err != nil {
				c.Error(err)
			}
			t2 := time.Now()

			// We reload the logger before we emit the log line so that we pick up any changes/additional
			// fields that might've been added in a downstream middleware.
			log = logger.FromContext(c.Request().Context())

			log.Root(logger.Data{
				"status_code": c.Response().Status,
				"duration":    fmt.Sprintf("%.5f", t2.Sub(t1).Seconds()*1000),
				"referer":     c.Request().Referer(),
				"user_agent":  c.Request().UserAgent(),
			}).Info("request handled")
			return nil
		}
	}
}

// IDFromEchoContext returns the request ID from the given echo.Context. If
// there is no request ID, then this will just return the empty string.
func IDFromEchoContext(c echo.Context) string {
	id, ok := c.Get(echoIDKey).(string)
	if !ok {
		return ""
	}
	return id
}

// FromEchoContext returns a logger.Logger from the given echo.Context. It fetches the
// logger on the underlying context.Context.
func FromEchoContext(c echo.Context) logger.Logger {
	return logger.FromContext(c.Request().Context())
}
