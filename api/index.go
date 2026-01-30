package handler

import (
	"net/http"

	"zach-sikora-daycare/internal/config"
	"zach-sikora-daycare/internal/handler"
	"zach-sikora-daycare/internal/middleware"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func init() {
	cfg := config.Load()

	e = echo.New()
	e.HideBanner = true
	e.HidePort = true

	middleware.Setup(e, cfg)

	// For Vercel, we skip database since SQLite doesn't work well in serverless
	// Use Neon PostgreSQL if you need a database in production
	h := handler.New(cfg, nil)
	h.RegisterRoutes(e)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e.ServeHTTP(w, r)
}
