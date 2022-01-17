package routes

import (
	"net/http"
	"strconv"

	"github.com/jeongpope/example-msa/services/backend/model"
	"github.com/jeongpope/example-msa/services/backend/movie"
	"github.com/labstack/echo/v4"
)

func getMovies(c echo.Context) error {
	movies := movie.GetMovies()

	return c.JSON(http.StatusOK, movies)
}

func getMovie(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		//		c.Echo().Logger.Error(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, "invalid syntax")
	}

	movie, err := movie.GetMovie(id)
	if err != nil {
		c.Echo().Logger.Error(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, movie)
}

func createMovie(c echo.Context) error {
	m := model.Movie{}
	if err := c.Bind(&m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	movie, err := movie.CreateMovie(m)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, movie)
}

func updateMovie(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		c.Echo().Logger.Error(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, "invalid syntax")
	}

	m := model.Movie{}
	if err := c.Bind(&m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	movie, err := movie.UpdateMovie(id, m)
	if err != nil {
		c.Echo().Logger.Error(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, movie)
}

func deleteMovie(c echo.Context) error {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		c.Echo().Logger.Error(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, "invalid syntax")
	}

	err = movie.DeleteMovie(id)
	if err != nil {
		c.Echo().Logger.Error(err.Error())

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "success")
}
