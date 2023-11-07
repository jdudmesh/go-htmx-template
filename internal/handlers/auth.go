package handlers

import (
	"bytes"
	"context"
	"gohtmx/internal/model"
	"gohtmx/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type datastore interface {
	CurrentUser(ctx context.Context) (*model.User, error)
}

func AuthButtons(store datastore) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := store.CurrentUser(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(500, err)
		}
		buf := bytes.NewBuffer(nil)
		err = views.AuthButtons(user).Render(c.Request().Context(), buf)
		if err != nil {
			log.Warn("TODO: you need to implement this properly")
			return echo.NewHTTPError(500, err)
		}
		return c.Blob(200, "text/html; charset=utf-8", buf.Bytes())
	}
}
