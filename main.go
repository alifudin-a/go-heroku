package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "==> METHOD=${method}, STATUS=${status}, HOST=${host}, URI=${uri}, " +
			"ERROR=${error}, LATENCY_HUMAN=${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/", acceesible)
	e.GET("/ping", ping)

	e.Logger.Fatal(e.Start(":3030"))

}

func acceesible(c echo.Context) (err error) {

	message := map[string]interface{}{
		"message": "accessible",
	}

	return c.JSON(http.StatusOK, message)
}

func ping(c echo.Context) (err error) {

	message := map[string]interface{}{
		"message": "pong",
	}

	return c.JSON(http.StatusOK, message)
}
