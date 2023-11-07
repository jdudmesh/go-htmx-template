package handlers

import (
	"bytes"
	"context"

	"gohtmx/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func HomePage() echo.HandlerFunc {
	buf := bytes.NewBuffer(nil)
	err := views.HomePage("Demo - Home", nil).Render(context.Background(), buf)
	if err != nil {
		log.Warn("TODO: you need to implement this properly")
		log.Errorf("rendering index: %s", err)
	}
	return func(c echo.Context) error {
		return c.Blob(200, "text/html; charset=utf-8", buf.Bytes())
	}
}
