package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT must be set!")
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "==> METHOD=${method}, STATUS=${status}, HOST=${host}, URI=${uri}, " +
			"ERROR=${error}, LATENCY_HUMAN=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	e.Static("/", "static")
	e.GET("/ping", ping)
	e.GET("/env", testEnv)

	e.Logger.Fatal(e.Start(":" + port))

}

func testEnv(c echo.Context) (err error) {

	env := os.Getenv("ENV_TEST")
	if env != "" {
		log.Println("Env Loaded!", string(env))
	}

	return c.JSON(http.StatusOK, "OK")
}

func ping(c echo.Context) (err error) {

	message := map[string]interface{}{
		"message": "pong",
	}

	return c.JSON(http.StatusOK, message)
}
