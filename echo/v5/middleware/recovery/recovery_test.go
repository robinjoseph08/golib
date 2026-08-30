package recovery

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/robinjoseph08/golib/echo/v5/middleware/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecovery(t *testing.T) {
	e := echo.New()
	e.Use(logger.Middleware())
	e.Use(Middleware())

	e.GET("/error", func(c *echo.Context) error { panic(errors.New("error")) })
	e.GET("/string", func(c *echo.Context) error { panic("string") })
	e.GET("/int", func(c *echo.Context) error { panic(1) })

	paths := []string{"/error", "/string", "/int"}

	for _, path := range paths {
		req, err := http.NewRequest("GET", path, nil)
		require.Nil(t, err, "unexpected error when making new request")

		w := httptest.NewRecorder()

		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code, "incorrect recovered status code")
		assert.Contains(t, w.Body.String(), "Internal Server Error", "incorrect error message")
	}
}

func TestRecoveryRethrowsAbortHandler(t *testing.T) {
	e := echo.New()
	c := e.NewContext(
		httptest.NewRequest(http.MethodGet, "/", nil),
		httptest.NewRecorder(),
	)
	handler := Middleware()(func(c *echo.Context) error {
		panic(http.ErrAbortHandler)
	})

	assert.PanicsWithValue(t, http.ErrAbortHandler, func() {
		_ = handler(c)
	})
}
