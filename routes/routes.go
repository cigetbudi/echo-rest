package routes

import (
	"echo-rest/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello ini echo")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai)

	return e
}
