package logger

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/robinjoseph08/golib/echo/test"
	"github.com/robinjoseph08/golib/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMiddleware(t *testing.T) {
	e := echo.New()
	e.Use(Middleware())

	e.GET("/", func(c echo.Context) error {
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

func TestIDFromEchoContext(t *testing.T) {
	e := echo.New()
	e.Use(Middleware())

	e.GET("/", func(c echo.Context) error {
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
