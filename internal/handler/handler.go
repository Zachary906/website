package handler

import (
	"zach-sikora-daycare/internal/config"
	"zach-sikora-daycare/internal/database"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg *config.Config
	db  *database.DB
}

func New(cfg *config.Config, db *database.DB) *Handler {
	return &Handler{
		cfg: cfg,
		db:  db,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.Static("/static", "static")

	e.GET("/health", h.Health)
	e.GET("/", h.Home)
}
