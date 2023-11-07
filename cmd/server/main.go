package main

import (
	client "github.com/jdudmesh/gomon-client"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"gohtmx/cmd/boot"
	"gohtmx/internal/handlers"
	"gohtmx/internal/store"
)

type Reloader struct {
}

func (r *Reloader) Reload(data string) {
	log.Debugf("Reloading trigger: %s", data)
}

func main() {
	cfg, err := boot.LoadConfig()
	if err != nil {
		panic(err)
	}
	log.Infof("Starting UI server: %s", cfg.CurrentDirectory)

	store, err := store.NewSqlite(cfg.DatabaseURL, true)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer store.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/assets", cfg.StaticFileDir)

	r := &Reloader{}
	t, err := client.New(r, e.Logger)
	if err != nil {
		log.Fatalf("unable to start reloader: %v", err)
	}
	defer t.Close()
	if err := t.Run(); err != nil {
		panic(err)
	}

	e.GET("/", handlers.HomePage())
	e.GET("/components/auth-buttons", handlers.AuthButtons(store))

	log.Fatal(e.Start(cfg.ServerAddress()))
}
