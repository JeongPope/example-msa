package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRoutes_GetMovies(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/movies", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getMovies(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestRoutes_GetMovie(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/movie/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, getMovie(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestRoutes_PostMovie(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/movie", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, createMovie(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestRoutes_PutMovie(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/api/movie/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, updateMovie(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestRoutes_DeleteMovie(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/movie/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, deleteMovie(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
