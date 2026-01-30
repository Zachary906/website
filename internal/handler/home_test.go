package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"zach-sikora-daycare/internal/testutil"

	"github.com/labstack/echo/v4"
)

func TestHealth(t *testing.T) {
	t.Parallel()

	db := testutil.NewTestDB(t)
	cfg := testutil.NewTestConfig(t)
	h := New(cfg, db)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.Health(c); err != nil {
		t.Errorf("Health() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Health() status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestHome(t *testing.T) {
	t.Parallel()

	db := testutil.NewTestDB(t)
	cfg := testutil.NewTestConfig(t)
	h := New(cfg, db)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.Home(c); err != nil {
		t.Errorf("Home() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Home() status = %d, want %d", rec.Code, http.StatusOK)
	}
}
