package recovery

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pkg/errors"
)

// Middleware recovers panics in subsequent handlers and funnels them to the
// error handler to be returned as a 500. It re-panics http.ErrAbortHandler so
// the server can abort the request.
func Middleware() func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					if r == http.ErrAbortHandler {
						panic(r)
					}
					if recoveredErr, ok := r.(error); ok {
						err = recoveredErr
					} else {
						err = errors.Errorf("%v", r)
					}
				}
			}()
			return next(c)
		}
	}
}
