package logger

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/robinjoseph08/golib/echo/v5/test"
	"github.com/robinjoseph08/golib/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMiddleware(t *testing.T) {
	e := echo.New()
	e.Use(Middleware())

	e.GET("/", func(c *echo.Context) error {
		log := FromEchoContext(c)
		assert.NotEqual(t, log.GetID(), "")
		return nil
	})

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestMiddlewareHandlesErrors(t *testing.T) {
	var output bytes.Buffer
	previousOutput := logger.Output()
	logger.SetOutput(&output)
	t.Cleanup(func() {
		logger.SetOutput(previousOutput)
	})

	e := echo.New()
	errorHandlerCalls := 0
	e.HTTPErrorHandler = func(c *echo.Context, err error) {
		errorHandlerCalls++
		require.EqualError(t, err, "boom")
		require.NoError(t, c.NoContent(http.StatusTeapot))
	}
	e.Use(Middleware())
	e.GET("/", func(c *echo.Context) error {
		return errors.New("boom")
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusTeapot, rr.Code)
	assert.Equal(t, 1, errorHandlerCalls)

	var logEntry map[string]interface{}
	require.NoError(t, json.Unmarshal(bytes.TrimSpace(output.Bytes()), &logEntry))
	assert.Equal(t, float64(http.StatusTeapot), logEntry["status_code"])
}

func TestIDFromEchoContext(t *testing.T) {
	e := echo.New()
	e.Use(Middleware())

	e.GET("/", func(c *echo.Context) error {
		id := IDFromEchoContext(c)
		assert.NotEqual(t, id, "")
		return nil
	})

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestFromEchoContext(t *testing.T) {
	log := logger.New().ID("foo")
	c, _ := test.NewContext(t, nil)
	c.SetRequest(c.Request().WithContext(log.WithContext(c.Request().Context())))

	l := FromEchoContext(c)

	assert.Equal(t, log.GetID(), l.GetID())

	c, _ = test.NewContext(t, nil)

	l = FromEchoContext(c)

	assert.Equal(t, "", l.GetID())
}
