package routes

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	port = ":"
)

func init() {
	envPort := os.Getenv("BACKEND_PORT")

	port += envPort
}

func Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Kubernetes bootcamp - Server is Up and Running...")
	})

	// All movies
	e.GET("/api/movies", getMovies)

	// Specific movie
	e.GET("/api/movie/:id", getMovie)
	e.POST("/api/movie", createMovie)
	e.PUT("/api/movie/:id", updateMovie)
	e.DELETE("/api/movie/:id", deleteMovie)

	e.Start(port)
}
